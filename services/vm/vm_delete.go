package vm

import (
	"fmt"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s VMService) Delete(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_DELETE, fmt.Sprintf("%s/%s", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_NO_CONTENT.Integer())
	if err != nil {
		return
	}

	success = true
	return
}
