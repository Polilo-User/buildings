package repository

import (
	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/services/favorites/model"
)

func SetFavorites(repo *repository, request model.SetFavoritesRequest) (err error) {
	req := `insert into "favorites" ("user_id","room_id") values ($1,$2)`
	stmt2, err := repo.db.Prepare(req)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	_, err = stmt2.Exec(request.UserId, request.RoomId) // Запускаем инсерт
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	defer stmt2.Close()
	return nil
}

func DeleteFavorites(repo *repository, request model.SetFavoritesRequest) (err error) {
	req := `delete from "favorites"  WHERE "user_id" = $1 AND "room_id" = $2 `
	stmt2, err := repo.db.Prepare(req)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	_, err = stmt2.Exec(request.UserId, request.RoomId) // Запускаем инсерт
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	defer stmt2.Close()
	return nil
}
