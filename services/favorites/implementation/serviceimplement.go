package implementation

import (
	"context"

	//"mime/multipart"

	"github.com/Polilo-User/buildings/functions/logging"
	favoritessvc "github.com/Polilo-User/buildings/services/favorites"

	"github.com/Polilo-User/buildings/services/favorites/model"
)

type service struct {
	repository favoritessvc.FavoritesRepo // репозитарий
	logger     *logging.Logger            // логгер
}

func NewService(rep favoritessvc.FavoritesRepo, logger *logging.Logger) favoritessvc.FavoritesService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) GetFavorites(ctx context.Context, request model.GetFavoritesRequest) (*model.GetFavoritesResponse, error) {
	return s.repository.GetFavorites(ctx, request)
}

func (s *service) SetFavorites(ctx context.Context, request model.SetFavoritesRequest) error {
	if request.Favorite {
		return s.repository.SetFavorites(ctx, request)
	} else {
		return s.repository.DeleteFavorite(ctx, request)
	}
}
