package models

type RatingModels struct {
	RatingID          string `json:"rating_id"`
	StoreID           string `json:"store_id"`
	UserID            string `json:"user_id"`
	RatingValue       string `json:"rating_value"`
	RatingDescription string `json:"rating_description"`
	RatingCreated     string `json:"rating_created"`
}
