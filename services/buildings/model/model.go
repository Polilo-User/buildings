package model

type Filters struct {
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
