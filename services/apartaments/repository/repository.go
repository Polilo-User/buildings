package repository

import (
	"context"
	"database/sql"

	"github.com/Polilo-User/buildings/functions/logging"
	apartSvc "github.com/Polilo-User/buildings/services/apartaments"
	"github.com/Polilo-User/buildings/services/apartaments/model"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) apartSvc.ApartService {
	return &repository{
		db:     db,
		logger: logger,
	}
}
func (repo *repository) GetApartByFilter(ctx context.Context, request model.GetApartByFilterRequest) (*model.GetApartByFilterResponse, error) {
	filters := getFilters(request.Filter)
	return GetApartamentsByFilter(repo, filters)
}
