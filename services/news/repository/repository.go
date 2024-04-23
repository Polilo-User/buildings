package repository

import (
	"context"
	"database/sql"

	"github.com/Polilo-User/buildings/functions/logging"
	newssvc "github.com/Polilo-User/buildings/services/news"
	"github.com/Polilo-User/buildings/services/news/model"
)

// Структура репозитария
type repository struct {
	db     *sql.DB         // БД
	logger *logging.Logger // Логгер
}

// Возвращает новый репозитарий с подключением к БД
func New(db *sql.DB, logger *logging.Logger) newssvc.NewsRepo {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (repo *repository) GetNews(ctx context.Context, request model.GetNewsRequest) (*model.GetNewsResponse, error) {
	return nil, nil
}
