package repository

import (
	"encoding/json"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/news/model"
)

func GetNews(repo *repository) (*model.GetNewsResponse, error) {
	var news []model.News
	req := "SELECT id, \"name\", \"imgUrl\", \"dtCreate\" FROM news"
	newsData, err := functions.Query2(repo.db, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	if len(newsData) == 0 {
		return nil, errors.NotFound.New("не смогли найти данные в БД")
	}
	// Парсим в структуру данные об установке очистки
	newsJson, err := json.Marshal(newsData)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	err = json.Unmarshal(newsJson, &news)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res := &model.GetNewsResponse{
		Data: news,
	}

	return res, nil
}
