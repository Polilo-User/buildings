package transport

import (
	"context"

	"github.com/Polilo-User/buildings/services/apartaments/model"

	apartaments "github.com/Polilo-User/buildings/services/apartaments"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints Структура со всеми эндпоинтами сервиса
type Endpoints struct {
	ApartByFilter endpoint.Endpoint
}

func MakeEndpoints(s apartaments.ApartService) Endpoints {
	return Endpoints{
		ApartByFilter: makeGetApartByFilter(s),
	}
}

func makeGetApartByFilter(s apartaments.ApartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetApartByFilter(ctx, request.(model.GetApartByFilterRequest))
	}
}
