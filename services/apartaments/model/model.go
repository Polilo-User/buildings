package model

type Apart struct {
	Name string `json:"name"`
}

type Filters struct {
}

type GetApartByFilterRequest struct {
	Filter Filters `json:"filter"`
}
type GetApartByFilterResponse struct {
	Data []Apart `json:"data"`
}

type LockApartRequest struct {
	ApartId int32 `json:"apart_id"`
}

type LockApartResponse struct {
	ApartId int32 `json:"apart_id"`
}
