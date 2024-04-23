package repository

import (
	"encoding/json"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/apartaments/model"
)

func GetApartamentsByFilter(repo *repository, filters string) (*model.GetApartByFilterResponse, error) {
	var apartaments []model.Apart
	req := "SELECT id, \"name\", \"price\" FROM apartaments" + filters
	apartamentsData, err := functions.Query2(repo.db, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	if len(apartamentsData) == 0 {
		return nil, errors.NotFound.New("не смогли найти данные в БД")
	}
	// Парсим в структуру данные об установке очистки
	apartamentsJson, err := json.Marshal(apartamentsData)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	err = json.Unmarshal(apartamentsJson, &apartaments)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res := &model.GetApartByFilterResponse{
		Data: apartaments,
	}

	return res, nil
}

func getFilters(model.Filters) string {
	return ""
}
