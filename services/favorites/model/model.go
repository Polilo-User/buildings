package model

type Apartaments struct {
}

type GetFavoritesRequest struct {
	Req string `json:"req"`
}
type GetFavoritesResponse struct {
	Data []Apartaments `json:"data"`
}
