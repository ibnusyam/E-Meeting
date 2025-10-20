package utils

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{Message: message, Data: data}
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Message: message}
}
