package transport

import (
	"context"

	news "github.com/Polilo-User/buildings/services/news"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints Структура со всеми эндпоинтами сервиса
type Endpoints struct {
	GetNews endpoint.Endpoint
}

func MakeEndpoints(s news.NewsService) Endpoints {
	return Endpoints{
		GetNews: makeGetBuildingsByFilter(s),
	}
}

func makeGetBuildingsByFilter(s news.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetNews(ctx)
	}
}
