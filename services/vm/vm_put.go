package vm

import (
	"fmt"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/vm_request"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s VMService) Update(uuid string, data vm_request.VMUpdateData) (success bool, err error) {
	serverRoute := fmt.Sprintf("%s/%s", constants.ROUTE_VM, uuid)
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_PUT, serverRoute, data)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_NO_CONTENT.Integer())
	if err != nil {
		return false, err
	}

	success = true
	return
}
