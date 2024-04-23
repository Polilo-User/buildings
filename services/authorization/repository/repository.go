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

func (repo *repository) GetStuffByUsernameAndPassord(ctx context.Context, username, password string) (string, string, string, string, string, string, []string, error) {
	roles := []string{}
	req := functions.AdaptReq("SELECT sm.id, sm.status, sm.\"isActive\" , sm.\"jobTitle\", f.url, smr.\"roleSignatura\", sm.fio FROM su_members sm inner join su_member2roles smr on sm.id = smr.\"memberId\" left join files f on sm.\"stuffPhoto\" = f.id WHERE sm.\"userName\"=? and sm.\"password\"=?")
	stuffData, err := functions.Query2(repo.db, req, username, password)
	if err != nil {
		return "", "", "", "", "", "", nil, errors.InternalServer.Wrap(err)
	}
	if len(stuffData) == 0 {
		return "", "", "", "", "", "", nil, errors.Teapot.NewCtx("не смогли найти пользователя с таким логином/паролем", "username:%v", username)
	}
	stuffId := functions.Intf2str(stuffData[0]["id"])
	status := functions.Intf2str(stuffData[0]["status"])
	isActive := functions.Intf2str(stuffData[0]["isActive"])
	fio := functions.Intf2str(stuffData[0]["fio"])
	jobTitle := functions.Intf2str(stuffData[0]["jobTitle"])
	stuffPhoto := functions.Intf2str(stuffData[0]["url"])
	for _, val := range stuffData {
		roles = append(roles, functions.Intf2str(val["roleSignatura"]))
	}
	return stuffId, status, isActive, fio, stuffPhoto, jobTitle, roles, nil
}

func (repo *repository) SetSession(ctx context.Context, stuffId string, session authsvc.Session) error {
	data, err := functions.Query2(repo.db, functions.AdaptReq("SELECT * FROM \"su_sessionTokens\" s WHERE s.\"memberId\"=?"), stuffId)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	if len(data) == 0 {
		req := functions.AdaptReq("insert into \"su_sessionTokens\" (\"memberId\",\"code\",\"updateDt\") values (?,?,?)")
		stmt2, err := repo.db.Prepare(req)
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		_, err = stmt2.Exec(stuffId, session.RefreshToken, session.UpdateDt) // Запускаем инсерт
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		defer stmt2.Close()
	} else {
		req := functions.AdaptReq("update \"su_sessionTokens\" set \"code\"=?,\"updateDt\"=? where \"memberId\"=?")
		stmt2, err := repo.db.Prepare(req)
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		_, err = stmt2.Exec(session.RefreshToken, session.UpdateDt, stuffId) // Запускаем инсерт
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		defer stmt2.Close()
	}
	return nil
}

func (repo *repository) GetStuffByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	stuffData, err := functions.Query2(repo.db, functions.AdaptReq("SELECT * FROM su_members s INNER JOIN \"su_sessionTokens\" st on s.\"id\"=st.\"memberId\" WHERE st.\"code\"=?"), refreshToken)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	if len(stuffData) == 0 {
		return "", errors.NotFound.New("Не смогли найти данные пользователя")
	}
	stuffId := functions.Intf2str(stuffData[0]["memberId"])
	expDateStr := functions.Intf2str(stuffData[0]["updateDt"])
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
	req := functions.AdaptReq("UPDATE su_members SET \"password\" = ?, \"status\" = null where \"id\" = ?")
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
