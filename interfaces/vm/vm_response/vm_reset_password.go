package vm_response

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/base_response"

type VMResetPasswordResponseData struct {
	base_response.BaseResponse
	Data vmResetPasswordData `json:"data"`
}

type vmResetPasswordData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
