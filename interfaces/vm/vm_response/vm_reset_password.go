package vm_response

type VMResetPasswordResponseData struct {
	Status int                 `json:"status"`
	Data   vmResetPasswordData `json:"data"`
}

type vmResetPasswordData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
