package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Polilo-User/buildings/functions/errors"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/logging"
	authsvc "github.com/Polilo-User/buildings/services/authorization"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) authsvc.AuthRepo {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (repo *repository) GetStuffByPhone(ctx context.Context, phoneNumber int, password string) (string, string, error) {
	req := `SELECT u.id, u."fullName" FROM users u WHERE u."phoneNumber" = $1 AND u.password = $2`
	stuffData, err := functions.Query2(repo.db, req, phoneNumber, password)
	if err != nil {
		return "", "", errors.InternalServer.Wrap(err)
	}
	if len(stuffData) == 0 {
		return "", "", errors.Teapot.NewCtx("не смогли найти пользователя с таким номером телефона", "username:%v", phoneNumber)
	}
	stuffId := functions.Intf2str(stuffData[0]["id"])
	fullName := functions.Intf2str(stuffData[0]["fullName"])

	return stuffId, fullName, nil
}

func (repo *repository) SetSession(ctx context.Context, stuffId string, session authsvc.Session) error {
	data, err := functions.Query2(repo.db, `SELECT * FROM "sessiontokens" s WHERE s."user_id"=$1`, stuffId)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	if len(data) == 0 {
		req := functions.AdaptReq(`insert into "sessiontokens" ("user_id","code","updated") values ($1,$2,now())`)
		stmt2, err := repo.db.Prepare(req)
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		_, err = stmt2.Exec(stuffId, session.RefreshToken) // Запускаем инсерт
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		defer stmt2.Close()
	} else {
		req := functions.AdaptReq(`update "sessiontokens" set \"code\"=$1,"updated"=now() where "user_id"=$2`)
		stmt2, err := repo.db.Prepare(req)
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		_, err = stmt2.Exec(session.RefreshToken, stuffId) // Запускаем инсерт
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		defer stmt2.Close()
	}
	return nil
}

func (repo *repository) GetStuffByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	stuffData, err := functions.Query2(repo.db, `SELECT * FROM users u INNER JOIN "sessiontokens" st on u."id"=u."userId" WHERE st."code"=$1`, refreshToken)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	if len(stuffData) == 0 {
		return "", errors.NotFound.New("Не смогли найти данные пользователя")
	}
	stuffId := functions.Intf2str(stuffData[0]["user_id"])
	expDateStr := functions.Intf2str(stuffData[0]["updated"])
	expDate, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", expDateStr)
	//expDate, err := time.ParseInLocation("2006-01-02 15:04:05", expDateStr, time.Local)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	if expDate.Before(time.Now()) {
		return "", errors.Teapot.New("Токен просрочен, пожалуйста авторизуйтесь заново.")
	}
	return stuffId, nil
}

// Изменить пароль пользователя
func (repo *repository) ChangePassword(ctx context.Context, stuffId, password string) error {
	req := "UPDATE users SET \"password\" = $1 where \"id\" = $2"
	stmt2, err := repo.db.Prepare(req)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	_, err = stmt2.Exec(password, stuffId) // Запускаем апдейт
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	return nil
}
