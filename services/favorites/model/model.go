package model

type Apartaments struct {
}

type GetFavoritesRequest struct {
	Req string `json:"req"`
}
type GetFavoritesResponse struct {
	Data []Apartaments `json:"data"`
}

type SetFavoritesRequest struct {
	UserId   int  `json:"user_id"`
	RoomId   int  `json:"room_id"`
	Favorite bool `json:"favorite"`
}
