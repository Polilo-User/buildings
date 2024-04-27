package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/buildings/model"
)

func GetBuildingsByFilter(repo *repository, filters string) (*model.GetBuildingsByFilterResponse, error) {
	var buildings []model.Buildings
	req := "SELECT id, \"name\", \"imgUrl\" FROM buildings b " + filters
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

func getFilters(filters model.Filters) (res string) {
	count := 0
	if filters.PriceFrom != 0 || filters.PriceTo != 0 {
		res += fmt.Sprintf("INNER JOIN rooms r ON b.id = r.building_id WHERE price > %d and price < %d", filters.PriceFrom, filters.PriceTo)
		count += 1
	}
	if filters.PassDate != "" {
		if count == 0 {
			res += "WHERE passDt = " + filters.PassDate
		} else {
			res += "AND passDt = " + filters.PassDate
		}
		count += 1
	}
	return res
}
