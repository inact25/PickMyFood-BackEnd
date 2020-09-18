package models

type RatingModels struct {
	RatingID          string `json:"ratingID"`
	StoreID           string `json:"storeID"`
	UserID            string `json:"userID"`
	RatingValue       string `json:"ratingValue"`
	RatingDescription string `json:"ratingDescription"`
	RatingCreated     string `json:"ratingCreated"`
	UserFirstname     string `json:"userFirstname"`
	UserLastname      string `json:"userLastname"`
}
