package model

type News struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

type GetNewsResponse struct {
	Data []News `json:"data"`
}
