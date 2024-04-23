package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Polilo-User/buildings/functions/errors"

	"github.com/Polilo-User/buildings/functions/logging"
	response "github.com/Polilo-User/buildings/functions/middleware"
	val "github.com/Polilo-User/buildings/functions/validator"
	"github.com/Polilo-User/buildings/services/favorites/model"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadReq = errors.BadRequest.New("Не хватает параметров или неверный URL")
)

func NewService(svcEndpoints Endpoints, logger *logging.Logger) http.Handler {
	// Настроим роутер и инициализируем http эндпоинты
	var (
		r            = mux.NewRouter()
		errorEncoder = kithttp.ServerErrorEncoder(response.EncodeMyErrorResponse)
	)
	options := []kithttp.ServerOption{
		errorEncoder,
	}
	r.Methods("POST").Path("/favorites/getFavorites").Handler(kithttp.NewServer(
		svcEndpoints.GetFavorites,
		decodeGetFavoritesRequest,
		response.EncodeResponse,
		options...,
	))

	return r
}

func decodeGetFavoritesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var req model.GetFavoritesRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}
	err = val.ValidateHttpReq(req, "")
	if err != nil {
		return nil, errors.UnprocessableEntity.NewCtx("Отсутствуют обязательные поля!", err.Error())
	}
	return req, nil
}
