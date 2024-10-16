package vm

import (
	"encoding/json"
	"fmt"

	"raznar.id/vm-control-hypervisor-wrapper/interfaces/vm/vm_response"
	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s VMService) GetDetails(uuid string) (resData *vm_response.VMDetailsResponseData, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_GET, fmt.Sprintf("%s/%s", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != constants.HTTP_STATUS_OK.Integer() {
		// Handle non-200 responses
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	resData = &vm_response.VMDetailsResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	return
}

func (s VMService) GetMetrics(uuid string) (resData *vm_response.VMMetricsResponseData, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_GET, fmt.Sprintf("%s/%s", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != constants.HTTP_STATUS_OK.Integer() {
		// Handle non-200 responses
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	resData = &vm_response.VMMetricsResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	return
}
