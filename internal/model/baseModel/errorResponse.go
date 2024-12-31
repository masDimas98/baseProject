package baseModel

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
