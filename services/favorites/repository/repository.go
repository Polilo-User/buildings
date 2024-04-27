package repository

import (
	"context"
	"database/sql"

	"github.com/Polilo-User/buildings/functions/logging"
	favoritessvc "github.com/Polilo-User/buildings/services/favorites"
	"github.com/Polilo-User/buildings/services/favorites/model"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) favoritessvc.FavoritesRepo {
	return &repository{
		db:     db,
		logger: logger,
	}
}
func (repo *repository) GetFavorites(ctx context.Context, request model.GetFavoritesRequest) (*model.GetFavoritesResponse, error) {
	return GetFavorites(repo)
}

func (repo *repository) SetFavorites(ctx context.Context, request model.SetFavoritesRequest) error {
	return SetFavorites(repo, request)
}

func (repo *repository) DeleteFavorite(ctx context.Context, request model.SetFavoritesRequest) error {
	return DeleteFavorites(repo, request)
}
