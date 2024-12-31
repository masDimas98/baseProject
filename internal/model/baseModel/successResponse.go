package baseModel

type SuccessResponse struct {
	Status      string `json:"status"`
	SuccessCode int    `json:"success_code"`
	Data        any    `json:"data"`
}
