package model

type Apart struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Area         int64  `json:"area"`
	Floor        int64  `json:"floor"`
	CountOfRooms int64  `json:"countOfRooms"`
}

type Filters struct {
	Area         int64 `json:"area"`
	CountOfRooms int64 `json:"rooms"`
	Floor        int64 `json:"floor"`
	PriceFrom    int64 `json:"priceFrom"`
	PriceTo      int64 `json:"priceTo"`
}

type GetApartByFilterRequest struct {
	Filter Filters `json:"filters"`
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
