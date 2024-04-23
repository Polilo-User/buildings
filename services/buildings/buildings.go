package buildings

import (
	"context"

	"github.com/Polilo-User/buildings/services/buildings/model"
)

type BuildingsService interface {
	GetBuildingsByFilter(ctx context.Context, request model.GetBuildingsByFilterRequest) (*model.GetBuildingsByFilterResponse, error)
}

type BuildingsRepo interface {
	GetBuildingsByFilter(ctx context.Context, request model.GetBuildingsByFilterRequest) (*model.GetBuildingsByFilterResponse, error)
}
