package models

type FeedbackModels struct {
	FeedbackID      string `json:"feedbackID"`
	StoreID         string `json:"storeID"`
	FeedbackValue   string `json:"feedbackValue"`
	FeedbackCreated string `json:"feedbackCreated"`
	StoreName       string `json:"storeName"`
}
