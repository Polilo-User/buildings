package auth

import (
	"context"

	"github.com/Polilo-User/buildings/services/authorization/model"
)

type Session struct {
	RefreshToken string `json:"refreshToken"`
	CreateDt     string `json:"createDt"`
	UpdateDt     string `json:"updateDt"`
}

type AuthService interface {
	Login(context.Context, model.LoginReq) (model.Tokens, error)
	CreateSession(ctx context.Context, StuffId string) (model.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (model.Tokens, error)
	ChangePassword(ctx context.Context, req model.ChangePasswordRequest) error
}

type AuthRepo interface {
	SetSession(ctx context.Context, StuffId string, session Session) error
	GetStuffByPhone(ctx context.Context, phoneNumber int, password string) (string, string, error)
	GetStuffByRefreshToken(ctx context.Context, refreshToken string) (string, error)
	ChangePassword(ctx context.Context, stuffId, password string) error
}
