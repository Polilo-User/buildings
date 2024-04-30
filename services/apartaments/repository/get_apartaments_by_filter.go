package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/apartaments/model"
)

func GetApartamentsByFilter(repo *repository, filters string) (*model.GetApartByFilterResponse, error) {
	var apartaments []model.Apart
	req := `SELECT coalesce(id, 0) as id,  coalesce(name, '') as name, coalesce(price, 0) as price, coalesce(area, 0) as area, coalesce(floor, 0) as floor FROM rooms` + filters
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

func getFilters(filters model.Filters) (res string) {
	count := 0
	if filters.Area != 0 {
		res += fmt.Sprintf(" WHERE area = %d", filters.Area)
		count += 1
	}
	if filters.CountOfRooms != 0 {
		if count == 0 {
			res += fmt.Sprintf(" WHERE countOfRooms = %d", filters.CountOfRooms)
			count += 1
		} else {
			res += fmt.Sprintf(" AND countOfRooms = %d", filters.CountOfRooms)
			count += 1
		}
	}
	if filters.Floor != 0 {
		if count == 0 {
			res += fmt.Sprintf(" WHERE floor = %d", filters.Floor)
			count += 1
		} else {
			res += fmt.Sprintf(" AND floor = %d", filters.Floor)
			count += 1
		}
	}

	if filters.PriceFrom != 0 || filters.PriceTo != 0 {
		if count == 0 {
			res += fmt.Sprintf(" WHERE price > %d and price < %d", filters.PriceFrom, filters.PriceTo)
			count += 1
		} else {
			res += fmt.Sprintf(" AND price > %d and price < %d", filters.PriceFrom, filters.PriceTo)
			count += 1
		}
	}

	return res
}
