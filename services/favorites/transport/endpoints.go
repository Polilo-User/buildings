package transport

import (
	"context"

	"github.com/Polilo-User/buildings/services/favorites"
	"github.com/Polilo-User/buildings/services/favorites/model"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints Структура со всеми эндпоинтами сервиса
type Endpoints struct {
	GetFavorites endpoint.Endpoint
}

func MakeEndpoints(s favorites.FavoritesService) Endpoints {
	return Endpoints{
		GetFavorites: makeGetFavorites(s),
	}
}

func makeGetFavorites(s favorites.FavoritesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetFavorites(ctx, request.(model.GetFavoritesRequest))
	}
}
