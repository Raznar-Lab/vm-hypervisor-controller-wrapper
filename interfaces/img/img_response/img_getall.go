package img_response

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/base_response"

type IMGGetAllResponseData struct {
	base_response.BaseResponse
	Data []string `json:"data"`
}
