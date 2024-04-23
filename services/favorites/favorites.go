package favorites

import (
	"context"

	"github.com/Polilo-User/buildings/services/favorites/model"
)

type FavoritesService interface {
	GetFavorites(ctx context.Context, request model.GetFavoritesRequest) (*model.GetFavoritesResponse, error)
}

type FavoritesRepo interface {
	GetFavorites(ctx context.Context, request model.GetFavoritesRequest) (*model.GetFavoritesResponse, error)
}
