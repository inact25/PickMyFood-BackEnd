package responses

type ResponsesData struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponsesStatus struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
