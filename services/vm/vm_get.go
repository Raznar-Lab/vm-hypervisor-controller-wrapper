package vm

import (
	"encoding/json"
	"fmt"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/dev/interfaces/vm/vm_response"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/dev/pkg/constants"
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

	resData = &vm_response.VMDetailsResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	err = s.HandleErrorResponse(&resData.BaseResponse, constants.HTTP_STATUS_OK.Integer())
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

	resData = &vm_response.VMMetricsResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	err = s.HandleErrorResponse(&resData.BaseResponse, constants.HTTP_STATUS_OK.Integer())
	return
}
