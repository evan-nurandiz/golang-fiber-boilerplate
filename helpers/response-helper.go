package helpers

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type EmptyObj struct{}

func BuildResponse(status int, message string, data interface{}) SuccessResponse {
	res := SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(status int, message string, err string) ErrorResponse {
	res := ErrorResponse{
		Status:  status,
		Message: message,
		Errors:  err,
	}
	return res
}
