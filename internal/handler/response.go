package handler

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type TaskResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
