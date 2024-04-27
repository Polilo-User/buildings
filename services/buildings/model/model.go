package model

type Filters struct {
	Cvartal   int   `json:"cvartal"`
	Year      int   `json:"year"`
	PriceFrom int64 `json:"priceFrom"`
	PriceTo   int64 `json:"priceTo"`
}

type Buildings struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

type GetBuildingsByFilterRequest struct {
	Filter Filters `json:"filters"`
}
type GetBuildingsByFilterResponse struct {
	Data []Buildings `json:"data"`
}
