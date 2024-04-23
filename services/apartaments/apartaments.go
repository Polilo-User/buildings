package apartaments

import (
	"context"

	"github.com/Polilo-User/buildings/services/apartaments/model"
)

type ApartService interface {
	GetApartByFilter(ctx context.Context, request model.GetApartByFilterRequest) (*model.GetApartByFilterResponse, error)
}

type ApartRepo interface {
	GetApartByFilter(ctx context.Context, request model.GetApartByFilterRequest) (*model.GetApartByFilterResponse, error)
}
