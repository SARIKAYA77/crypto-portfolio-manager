package model

type Currency struct {
	ID         int     `json:"id"`
	Code       string  `json:"code"`
	Amount     float64 `json:"amount"`
	Price      float64 `json:"price"`
	TotalValue float64 `json:"totalValue"`
	History    []Price `json:"history"`
}
type Price struct {
	Amount float64 `json:"amount"`
	Price  struct {
		Old     float64 `json:"old"`
		Current float64 `json:"current"`
	} `json:"price"`
}
