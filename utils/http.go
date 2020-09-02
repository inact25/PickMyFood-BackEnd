package utils

import (
	"encoding/json"
	"net/http"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/responses"
)

const (
	RESPONSE_SUCCESS  = "Success"
	BAD_REQUEST       = "Bad Request"
	BAD_GATEWAY       = "Bad Gateway"
	NOT_FOUND         = "Not Found"
	STATUS_NO_CONTENT = "No Content"
	STATUS_CREATED    = "Created"
)

func messageStatusCode(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return BAD_REQUEST
	case http.StatusBadGateway:
		return BAD_GATEWAY
	case http.StatusNotFound:
		return NOT_FOUND
	case http.StatusOK:
		return RESPONSE_SUCCESS
	case http.StatusCreated:
		return STATUS_CREATED
	default:
		return BAD_REQUEST
	}
}

func HandleRequest(res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := responses.ResponsesStatus{statusCode, messageStatusCode(statusCode)}
	byteOfResponseInsertion, _ := json.Marshal(response)
	res.Write(byteOfResponseInsertion)
}

func HandleResponseError(res http.ResponseWriter, statusCode int, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := responses.ResponsesStatus{statusCode, message}
	byteOfResponseInsertion, _ := json.Marshal(response)
	res.Write(byteOfResponseInsertion)
}

func HandleResponse(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := responses.ResponsesData{statusCode, messageStatusCode(statusCode), data}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}
