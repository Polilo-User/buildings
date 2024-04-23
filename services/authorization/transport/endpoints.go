package transport

import (
	"context"

	"github.com/Polilo-User/buildings/services/authorization/model"

	auth "github.com/Polilo-User/buildings/services/authorization"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints Структура со всеми эндпоинтами сервиса
type Endpoints struct {
	Login        endpoint.Endpoint
	RefreshToken endpoint.Endpoint
	//CreateIssue  endpoint.Endpoint
	CreateDeal      endpoint.Endpoint
	ChangePassword  endpoint.Endpoint
	B24refreshToken endpoint.Endpoint
}

func MakeEndpoints(s auth.AuthService) Endpoints {
	return Endpoints{
		Login:          makeLoginEndpoint(s),
		RefreshToken:   makeRefreshTokenEndpoint(s),
		ChangePassword: makeChangePasswordEndpoint(s),
	}
}

// makeLoginEndpoint Инициализация эндпоита для логина
func makeLoginEndpoint(s auth.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Login(ctx, request.(model.LoginReq))
	}
}

func makeRefreshTokenEndpoint(s auth.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.RefreshTokens(ctx, request.(model.RefreshTokensRequest).RefreshToken)
	}
}

func makeChangePasswordEndpoint(s auth.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.ChangePassword(ctx, request.(model.ChangePasswordRequest))
		if err != nil {
			return nil, err
		}
		return model.ChangePasswordResponse{
			Success: true,
		}, nil
	}
}
