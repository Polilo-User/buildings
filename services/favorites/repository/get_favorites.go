package repository

import (
	"encoding/json"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/favorites/model"
)

func GetFavorites(repo *repository) (res *model.GetFavoritesResponse, err error) {
	var rooms []model.Apartaments
	req := `SELECT coalesce(id, 0), coalesce("room_id", ''), coalesce("user_id", '') FROM favorites`
	favoritesData, err := functions.Query2(repo.db, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	if len(favoritesData) == 0 {
		return nil, errors.NotFound.New("не смогли найти данные в БД")
	}
	// Парсим в структуру данные об установке очистки
	favoritesJson, err := json.Marshal(favoritesData)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	err = json.Unmarshal(favoritesJson, &rooms)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res = &model.GetFavoritesResponse{
		Data: rooms,
	}

	return res, nil
}
