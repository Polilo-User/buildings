package implementation

import (
	"context"

	//"mime/multipart"

	"github.com/Polilo-User/buildings/functions/logging"
	newssvc "github.com/Polilo-User/buildings/services/news"

	"github.com/Polilo-User/buildings/services/news/model"
)

type service struct {
	repository newssvc.NewsRepo // репозитарий
	logger     *logging.Logger  // логгер
}

func NewService(rep newssvc.NewsRepo, logger *logging.Logger) newssvc.NewsService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) GetNews(ctx context.Context) (*model.GetNewsResponse, error) {
	return s.repository.GetNews(ctx)
}
