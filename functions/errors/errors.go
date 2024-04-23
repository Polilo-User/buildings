package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// Определяем типы ошибок
const (
	NoType              = ErrorType(500)
	LogicError          = ErrorType(1)
	BadRequest          = ErrorType(http.StatusBadRequest)
	NotFound            = ErrorType(http.StatusNotFound)
	Teapot              = ErrorType(http.StatusTeapot)
	InternalServer      = ErrorType(http.StatusInternalServerError)
	Forbidden           = ErrorType(http.StatusForbidden)
	Unauthorized        = ErrorType(http.StatusUnauthorized)
	UnprocessableEntity = ErrorType(http.StatusUnprocessableEntity)
)

type ErrorType uint

type CustomError struct {
	errorType   ErrorType `json:"-"`
	HumanText   string    `json:"humanTextError"`
	DevelopText string    `json:"developerTextError"`
	err         error     `json:"-"`
	Path        string    `json:"path"`
	Context     string    `json:"context,omitempty"`
}

func (error CustomError) Error() string {
	return error.err.Error()
}

// Создаем новую ошибку
func (typ ErrorType) New(msg string) error {
	return typ.NewPathCtx(msg, 2, "")
}

// Создаем новую ошибку с контекстом
func (typ ErrorType) NewCtx(msg, context string, args ...any) error {
	return typ.NewPathCtx(msg, 2, context, args...)
}

// Создаем новую ошибку с выбором глубины пути (от 1)
func (typ ErrorType) NewPath(msg string, funcNum int) error {
	return typ.NewPathCtx(msg, funcNum+1, "")
}

// Создаем новую ошибку с выбором глубины пути (от 1) и контекстом
func (typ ErrorType) NewPathCtx(msg string, funcNum int, context string, args ...any) error {
	_, file, line, _ := runtime.Caller(funcNum)
	return CustomError{
		errorType: typ,
		err:       errors.New(msg),
		Path:      fmt.Sprintf("%v:%v", file, line),
		Context:   fmt.Sprintf(context, args...),
	}
}

// Оборачиваем дефолтную ошибку в кастомный тип
func (typ ErrorType) Wrap(err error) error {
	return typ.WrapPathCtx(err, 2, "")
}

// Оборачиваем дефолтную ошибку в кастомный тип и задаем ей контекст
func (typ ErrorType) WrapCtx(err error, context string, agrs ...any) error {
	return typ.WrapPathCtx(err, 2, context, agrs...)
}

// Оборачиваем дефолтную ошибку в кастомный тип с выбором глубины пути (от 1)
func (typ ErrorType) WrapPath(err error, funcNum int) error {
	return typ.WrapPathCtx(err, funcNum+1, "")
}

// Оборачиваем дефолтную ошибку в кастомный тип с выбором глубины пути (от 1) и контекстом
func (typ ErrorType) WrapPathCtx(err error, funcNum int, context string, args ...any) error {

	_, file, line, _ := runtime.Caller(funcNum)

	if customErr, ok := err.(CustomError); ok {

		return CustomError{
			errorType: customErr.errorType,
			err:       customErr.err,
			HumanText: customErr.HumanText,
			Path:      fmt.Sprintf("%v:%v", file, line),
			Context:   customErr.Context,
		}
	}

	return CustomError{
		errorType: typ,
		err:       err,
		Path:      fmt.Sprintf("%v:%v", file, line),
	}
}

// Добавляем в ошибку текст, который можно отдать пользователю
func AddHumanText(err error, message string) error {

	if customErr, ok := err.(CustomError); ok {

		if customErr.HumanText != "" {
			return err
		}

		return CustomError{
			errorType: customErr.errorType,
			err:       customErr.err,
			HumanText: message,
			Path:      customErr.Path,
			Context:   customErr.Context,
		}
	}

	_, file, line, _ := runtime.Caller(1)

	return CustomError{
		errorType: NoType,
		err:       err,
		HumanText: message,
		Path:      fmt.Sprintf("%v:%v", file, line),
	}
}

// Получаем тип ошибки
func GetType(err error) ErrorType {

	if customErr, ok := err.(CustomError); ok {
		return customErr.errorType
	}

	return NoType
}

// Переводим ошибку в JSON
func Json(err error) ([]byte, error) {

	if customErr, ok := err.(CustomError); ok {

		pathArr := strings.Split(customErr.Path, "github.com/Polilo-User/buildings/")
		if len(pathArr) > 1 {
			customErr.Path = pathArr[1]
		}

		customErr.DevelopText = customErr.err.Error()

		byt, e := json.Marshal(customErr)
		if e != nil {
			return nil, InternalServer.Wrap(e)
		}

		return byt, nil
	}

	return nil, InternalServer.NewCtx("Дефолтная ошибка не обернута. Ошибка: %v", err.Error())
}
