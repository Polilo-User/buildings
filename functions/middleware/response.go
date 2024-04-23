package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Polilo-User/buildings/functions/errors"
	"github.com/Polilo-User/buildings/functions/logging"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func EncodeMyErrorResponse(_ context.Context, err error, w http.ResponseWriter) {

	errorType := errors.GetType(err)
	if errorType == errors.NoType {
		err = errors.NoType.Wrap(err)
	}

	logger := logging.GetLogger()

	switch errorType {
	case errors.InternalServer:
		logger.Error(err)
	case errors.Forbidden:
		logger.Warning(err)
	default:
		logger.Warning(err)
	}

	switch errorType {
	case errors.BadRequest:
		err = errors.AddHumanText(err, "Введены неверные данные")
	case errors.InternalServer:
		err = errors.AddHumanText(err, "Произошла непредвиденная ошибка")
	case errors.NotFound:
		err = errors.AddHumanText(err, "Данные не найдены")
	case errors.Forbidden:
		err = errors.AddHumanText(err, "Доступ запрещен")
	case errors.Unauthorized:
		err = errors.AddHumanText(err, "Пользователь не авторизован")
	}

	if errorType == errors.NoType || errorType == errors.LogicError {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(errorType))

	byt, e := errors.Json(err)
	if e != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
	}

	w.Write(byt)
}
