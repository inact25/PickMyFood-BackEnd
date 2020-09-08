package models

type FeedBack struct {
	FeedBackID      string `json:"feedbackID"`
	FeedBackValue   string `json:"feedbackValue"`
	FeedBackCreated string `json:"feedbackCreated"`
	Store           Store  `json;"store"`
	User            User   `json:"user"`
}
