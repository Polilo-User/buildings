package transport

import (
	"context"

	"github.com/Polilo-User/buildings/services/buildings/model"

	buildings "github.com/Polilo-User/buildings/services/buildings"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints Структура со всеми эндпоинтами сервиса
type Endpoints struct {
	GetBuildingsByFilter endpoint.Endpoint
}

func MakeEndpoints(s buildings.BuildingsService) Endpoints {
	return Endpoints{
		GetBuildingsByFilter: makeGetBuildingsByFilter(s),
	}
}

func makeGetBuildingsByFilter(s buildings.BuildingsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetBuildingsByFilter(ctx, request.(model.GetBuildingsByFilterRequest))
	}
}
