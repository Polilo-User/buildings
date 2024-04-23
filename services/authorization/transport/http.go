package transport

import (
	"context"
	"encoding/json"

	//"mime/multipart"
	"net/http"

	"github.com/Polilo-User/buildings/functions"
	"github.com/Polilo-User/buildings/functions/errors"

	"github.com/Polilo-User/buildings/functions/logging"
	response "github.com/Polilo-User/buildings/functions/middleware"
	val "github.com/Polilo-User/buildings/functions/validator"
	"github.com/Polilo-User/buildings/services/authorization/model"
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

	r.Methods("POST").Path("/authorization/login").Handler(kithttp.NewServer(
		svcEndpoints.Login,
		decodeLoginRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/authorization/refreshToken").Handler(kithttp.NewServer(
		svcEndpoints.RefreshToken,
		decodeRefreshTokensRequest,
		response.EncodeResponse,
		options...,
	))
	// r.Methods("POST").Path("/authorization/createIssue").Handler(kithttp.NewServer(
	// 	svcEndpoints.CreateIssue,
	// 	decodeCreateIssueRequest,
	// 	response.EncodeResponse,
	// 	options...,
	// ))
	r.Methods("POST").Path("/authorization/createDeal").Handler(kithttp.NewServer(
		svcEndpoints.CreateDeal,
		decodeCreateDealRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/authorization/changePassword").Handler(kithttp.NewServer(
		svcEndpoints.ChangePassword,
		decodeChangePasswordRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/authorization/b24refreshToken").Handler(kithttp.NewServer(
		svcEndpoints.B24refreshToken,
		decodeB24refreshTokenRequest,
		response.EncodeResponse,
		options...,
	))
	return r
}

// Login godoc
// @Summary Логин
// @Description
// @Produce json
// @Accept  json
// @Tags authorization
// @Param params body model.LoginReq true "Json params"
// @Param Custom-User-Agent header string false "User-Agent"
// @Success 200 {object} model.Tokens
// @Router /authorization/login [post]
func decodeLoginRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req model.LoginReq
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

// RefreshToken godoc
// @Summary Обновление токенов
// @Produce json
// @Accept  json
// @Tags authorization
// @Param params body model.RefreshTokensRequest true "Json params"
// @Param Custom-User-Agent header string false "User-Agent"
// @Success 200 {object} model.Tokens
// @Router /authorization/refreshToken [post]
func decodeRefreshTokensRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req model.RefreshTokensRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, errors.BadRequest.Wrap(err)
	} else if len(req.RefreshToken) == 0 {
		return nil, ErrBadReq
	}
	return req, nil
}

// // CreateIssue godoc
// // @Summary Тестовая апи для работы с яндекс-трекером
// // @Produce json
// // @Accept  json
// // @Tags authorization
// // @Param params body model.CreateIssueReq true "Json params"
// // @Param Custom-User-Agent header string false "User-Agent"
// // @Success 200 {object} model.Issue
// // @Router /authorization/createIssue [post]
// func decodeCreateIssueRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
// 	issueId := r.FormValue("issueId")

// 	var files []*multipart.FileHeader
// 	if r.MultipartForm != nil {
// 		files = r.MultipartForm.File["files"]
// 		if files == nil {
// 			files = r.MultipartForm.File["files[]"]
// 		}
// 	}

// 	req := model.IssueAttachmentReq{
// 		IssueId: issueId,
// 		Files:   files,
// 	}

// 	err = val.ValidateHttpReq(req, "")
// 	if err != nil {
// 		return nil, errors.UnprocessableEntity.NewCtx("Отсутствуют обязательные поля!", err.Error())
// 	}
// 	return req, nil
// }

// CreateDeal godoc
// @Summary Тестовая апи для работы с битрикс24
// @Produce json
// @Accept  json
// @Tags authorization
// @Param params body model.CreateDealRequest true "Json params"
// @Param Custom-User-Agent header string false "User-Agent"
// @Success 200 {object} model.CreateDealResponse
// @Router /authorization/createDeal [post]
func decodeCreateDealRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req model.CreateDealRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, errors.BadRequest.Wrap(err)
	}
	return req, nil
}

// ChangePassword godoc
// @Summary Изменить пароль пользователя
// @Produce json
// @Accept  json
// @Tags authorization
// @Param params body model.ChangePasswordRequest true "Json params"
// @Param Custom-User-Agent header string false "User-Agent"
// @Success 200 {object} model.ChangePasswordResponse
// @Router /authorization/changePassword [post]
func decodeChangePasswordRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	stuffId, err := functions.BearerAuth(r)
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}
	var body model.ChangePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}
	req := model.ChangePasswordRequest{
		StuffId:  stuffId,
		Password: body.Password,
	}
	err = val.ValidateHttpReq(req, "")
	if err != nil {
		return nil, errors.UnprocessableEntity.NewCtx("Отсутствуют обязательные поля!", err.Error())
	}
	return req, nil
}

func decodeB24refreshTokenRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req model.B24refreshTokenRequest
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
