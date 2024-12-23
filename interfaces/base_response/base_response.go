package base_response

type BaseResponse struct {
	Code   int                `json:"code"`
	Errors *BaseErrorResponse `json:"errors"`
}
