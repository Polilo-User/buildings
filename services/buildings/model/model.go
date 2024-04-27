package model

type Filters struct {
	PassDate  string
	PriceFrom int64
	PriceTo   int64
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
