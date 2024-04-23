package implementation

import (
	"context"

	"github.com/Polilo-User/buildings/functions/logging"

	apartSvc "github.com/Polilo-User/buildings/services/apartaments"
	"github.com/Polilo-User/buildings/services/apartaments/model"
)

// Реализация сервиса счетчиков
type service struct {
	repository apartSvc.ApartRepo // репозитарий
	logger     *logging.Logger    // логгер
}

func NewService(rep apartSvc.ApartRepo, logger *logging.Logger) apartSvc.ApartService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}
func (s *service) GetApartByFilter(ctx context.Context, request model.GetApartByFilterRequest) (*model.GetApartByFilterResponse, error) {
	return s.repository.GetApartByFilter(ctx, request)
}
