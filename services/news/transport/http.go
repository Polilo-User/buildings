package transport

import (
	"context"
	"net/http"

	"github.com/Polilo-User/buildings/functions/errors"

	"github.com/Polilo-User/buildings/functions/logging"
	response "github.com/Polilo-User/buildings/functions/middleware"
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
	r.Methods("GET").Path("/news/getNews").Handler(kithttp.NewServer(
		svcEndpoints.GetNews,
		decodeGetNewsRequest,
		response.EncodeResponse,
		options...,
	))

	return r
}

func decodeGetNewsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}
