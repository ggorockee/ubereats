package dto

type SearchRestaurant struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
}
