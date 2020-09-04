package models

type FeedbackModels struct {
	FeedbackID      string `json:"feedback_id"`
	StoreID         string `json:"store_id"`
	FeedbackValue   string `json:"feedback_value"`
	FeedbackCreated string `json:"feedback_created"`
}
