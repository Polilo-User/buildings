package model

type Filters struct {
	Cvartal   int64 `json:"cvartal"`
	Year      int64 `json:"year"`
	PriceFrom int64 `json:"priceFrom"`
	PriceTo   int64 `json:"priceTo"`
}

type Buildings struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ImgUrl    string `json:"imgUrl"`
	Cvartal   string `json:"cvartal"`
	PriceFrom int    `json:"price"`
}

type GetBuildingsByFilterRequest struct {
	Filter Filters `json:"filters"`
}
type GetBuildingsByFilterResponse struct {
	Data []Buildings `json:"data"`
}
