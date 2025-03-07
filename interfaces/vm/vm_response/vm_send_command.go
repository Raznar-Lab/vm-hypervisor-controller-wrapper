package vm_response

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/base_response"

type VMSendCommandResponseData struct {
	base_response.BaseResponse
	Data vmSendCommandData `json:"data"`
}

type vmSendCommandData struct {
	ExitCode int    `json:"exit_code"`
	Exited   int    `json:"exited"`
	OutData  string `json:"out-data"`
}
