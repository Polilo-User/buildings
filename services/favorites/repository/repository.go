package repository

import (
	"context"
	"database/sql"

	"github.com/Polilo-User/buildings/functions/logging"
	"github.com/Polilo-User/buildings/services/buildings/model"
	favoritessvc "github.com/Polilo-User/buildings/services/favorites"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) favoritessvc.FavoritesService {
	return &repository{
		db:     db,
		logger: logger,
	}
}
func (repo *repository) GetFavorites(ctx context.Context, request model.GetFavoritesRequest) (*model.GetFavoritesResponse, error) {
	return nil, nil
}
