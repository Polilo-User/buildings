package repository

import (
	"encoding/json"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/buildings/model"
)

func GetBuildingsByFilter(repo *repository, filters string) (*model.GetBuildingsByFilterResponse, error) {
	var buildings []model.Buildings
	req := "SELECT id, \"name\", \"imgUrl\" FROM buildings" + filters
	buildingsData, err := functions.Query2(repo.db, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	if len(buildingsData) == 0 {
		return nil, errors.NotFound.New("не смогли найти данные в БД")
	}
	// Парсим в структуру данные об установке очистки
	buildingsJson, err := json.Marshal(buildingsData)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	err = json.Unmarshal(buildingsJson, &buildings)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res := &model.GetBuildingsByFilterResponse{
		Data: buildings,
	}

	return res, nil
}

func getFilters(model.Filters) string {
	return ""
}
