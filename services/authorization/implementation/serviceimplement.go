package implementation

import (
	"context"
	"time"

	"github.com/Polilo-User/buildings/functions"

	"github.com/Polilo-User/buildings/functions/logging"

	//yandextracker "github.com/Polilo-User/buildings/functions/external/yandexTracker"
	authsvc "github.com/Polilo-User/buildings/services/authorization"
	"github.com/Polilo-User/buildings/services/authorization/model"
)

// Реализация сервиса счетчиков
type service struct {
	repository authsvc.AuthRepo // репозитарий
	logger     *logging.Logger  // логгер
}

// Создаем и возвращаем новый сервис счетчиков
func NewService(rep authsvc.AuthRepo, logger *logging.Logger) authsvc.AuthService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Login(ctx context.Context, input model.LoginReq) (res model.Tokens, err error) {
	stuffId, status, isActive, fio, stuffPhoto, jobTitle, roles, err := s.repository.GetStuffByUsernameAndPassord(ctx, input.Username, input.Password)
	if err != nil {
		return res, err
	}
	tokens, err := s.CreateSession(ctx, stuffId)
	if err != nil {
		return res, err
	}
	tokens.Status, tokens.IsActive, tokens.Fio, tokens.StuffPhoto, tokens.JobTitle = status, isActive, fio, stuffPhoto, jobTitle
	tokens.Roles = roles
	return tokens, nil
}

func (s *service) CreateSession(ctx context.Context, stuffId string) (res model.Tokens, err error) {
	manag, err := functions.NewManager(functions.SECRET_KEY)
	if err != nil {
		return res, err
	}
	dur, err := time.ParseDuration(functions.ACCESS_TOKEN_TTL)
	if err != nil {
		return res, err
	}
	res.AccessToken, err = manag.NewJWT(stuffId, dur)
	if err != nil {
		return res, err
	}
	res.RefreshToken, err = manag.NewRefreshToken()
	if err != nil {
		return res, err
	}
	refreshDur, err := time.ParseDuration(functions.REFRESH_TOKEN_TTL)
	if err != nil {
		return res, err
	}
	session := authsvc.Session{
		RefreshToken: res.RefreshToken,
		CreateDt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdateDt:     time.Now().Add(refreshDur).Format("2006-01-02 15:04:05"),
	}
	err = s.repository.SetSession(ctx, stuffId, session)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) RefreshTokens(ctx context.Context, refreshToken string) (res model.Tokens, err error) {
	stuff, err := s.repository.GetStuffByRefreshToken(ctx, refreshToken)
	if err != nil {
		return res, err
	}
	return s.CreateSession(ctx, stuff)
}

// Изменить пароль пользователя
func (s *service) ChangePassword(ctx context.Context, req model.ChangePasswordRequest) (err error) {
	err = s.repository.ChangePassword(ctx, req.StuffId, req.Password)
	if err != nil {
		return err
	}
	return nil
}
