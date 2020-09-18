package models

type FeedbackModels struct {
	FeedbackID      string `json:"feedbackID"`
	StoreID         string `json:"storeID"`
	UserID          string `json:"userID"`
	FeedbackValue   string `json:"feedbackValue"`
	FeedbackCreated string `json:"feedbackCreated"`
	UserFirstName   string `json:"userFirstname"`
	UserLastName    string `json:"userLastname"`
}
