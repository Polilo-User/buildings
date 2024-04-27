package model

import (
	"encoding/json"
	//"mime/multipart"
)

type LoginReq struct {
	PhoneNumber int    `json:"phone" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Tokens Tokens
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Fio          string `json:"fio"`
}

type RefreshTokensRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Структуры запросов битрикса
// Создание сделки
type CreateDealRequest struct {
	Fields DealFields  `json:"fields"`
	Params *DealParams `json:"params"`
}
type DealFields struct {
	Title        string      `json:"TITLE"`
	TypeId       string      `json:"TYPE_ID"`
	StageId      string      `json:"STAGE_ID"`
	CompanyId    json.Number `json:"COMPANY_ID"`
	ContactId    json.Number `json:"CONTACT_ID"`
	Opened       string      `json:"OPENED"`
	AssignedById json.Number `json:"ASSIGNED_BY_ID"`
	Probability  json.Number `json:"PROBABILITY"`
	CurrencyId   string      `json:"CURRENCY_ID"`
	Opportunity  json.Number `json:"OPPORTUNITY"`
	CategoryId   json.Number `json:"CATEGORY_ID"`
	BeginsDate   string      `json:"beginsdate"`
	CloseDate    string      `json:"closedate"`
}
type DealParams struct {
	RegisterSonetEvent string `json:"register_sonet_event"`
}

type CreateDealResponse struct {
	Result json.Number `json:"result"`
	Time   DealTime    `json:"time"`
}
type DealTime struct {
	Start            json.Number `json:"start"`
	Finish           json.Number `json:"finish"`
	Duration         json.Number `json:"duration"`
	Processing       json.Number `json:"processing"`
	DateStart        string      `json:"date_start"`
	DateFinish       string      `json:"date_finish"`
	OperatingResetAt json.Number `json:"operating_reset_at"`
	Operating        json.Number `json:"operating"`
}

// Сменить пароль
type ChangePasswordRequest struct {
	StuffId  string `json:"-"`
	Password string `json:"password" binding:"required"`
}
type ChangePasswordResponse struct {
	Success bool `json:"success"`
}

// системная, для записи новых токенов
type B24refreshTokenRequest struct {
	ApiKey string `json:"apiKey"`
}
type B24refreshTokenResponse struct {
	Success bool `json:"success"`
}
