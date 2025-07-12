package models

type Pizza struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}
