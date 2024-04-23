package model

type News struct {
	Name string `json:"name"`
}

type GetNewsRequest struct{}

type GetNewsResponse struct {
	Data []News `json:"data"`
}
