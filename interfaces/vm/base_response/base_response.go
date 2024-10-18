package base_response

type BaseResponse struct {
	Status int                `json:"status"`
	Errors *BaseErrorResponse `json:"errors"`
}
