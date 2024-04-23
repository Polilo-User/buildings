package motoHour

import (
	"context"

	"github.com/Polilo-User/buildings/services/news/model"
)

type NewsService interface {
	GetNews(ctx context.Context, request model.GetNewsRequest) (*model.GetNewsResponse, error)
}

type NewsRepo interface {
	GetNews(ctx context.Context, request model.GetNewsRequest) (*model.GetNewsResponse, error)
}
