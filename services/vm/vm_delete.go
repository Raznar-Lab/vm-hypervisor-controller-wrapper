package vm

import (
	"fmt"
	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s VMService) Delete(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_DELETE, fmt.Sprintf("%s/%s", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if res.StatusCode != constants.HTTP_STATUS_NO_CONTENT.Integer() {
		if err != nil {
			return
		}

		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_NO_CONTENT, res.StatusCode, res.Status)
		return
	}

	defer res.Body.Close()

	success = true
	return
}