package repository

import (
	"context"
	"database/sql"

	"github.com/Polilo-User/buildings/functions/logging"
	buildingssvc "github.com/Polilo-User/buildings/services/buildings"
	"github.com/Polilo-User/buildings/services/buildings/model"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) buildingssvc.BuildingsRepo {
	return &repository{
		db:     db,
		logger: logger,
	}
}
func (repo *repository) GetBuildingsByFilter(ctx context.Context, request model.GetBuildingsByFilterRequest) (*model.GetBuildingsByFilterResponse, error) {
	filters := getFilters(request.Filter)
	return GetBuildingsByFilter(repo, filters)
}
