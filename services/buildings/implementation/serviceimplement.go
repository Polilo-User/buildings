package implementation

import (
	"context"

	//"mime/multipart"

	"github.com/Polilo-User/buildings/functions/logging"
	buildingssvc "github.com/Polilo-User/buildings/services/buildings"

	"github.com/Polilo-User/buildings/services/buildings/model"
)

type service struct {
	repository buildingssvc.BuildingsRepo // репозитарий
	logger     *logging.Logger            // логгер
}

func NewService(rep buildingssvc.BuildingsRepo, logger *logging.Logger) buildingssvc.BuildingsService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) GetBuildingsByFilter(ctx context.Context, request model.GetBuildingsByFilterRequest) (*model.GetBuildingsByFilterResponse, error) {
	return s.repository.GetBuildingsByFilter(ctx, request)
}
