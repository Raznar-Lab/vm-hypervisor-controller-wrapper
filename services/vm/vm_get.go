package vm

import (
	"encoding/json"
	"fmt"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/vm_response"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s VMService) GetInstallOS(uuid string) (resData *vm_response.VMInstallOSData, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_GET, fmt.Sprintf("%s/%s/install-os", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData = &vm_response.VMInstallOSData{}
	resData.Code = res.StatusCode
	if res.Body != nil {
		err = json.NewDecoder(res.Body).Decode(resData)
		if err != nil {
			resData = nil
			return
		}
	}

	return
}

func (s VMService) IsQMAgentReady(uuid string) (v bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_GET, fmt.Sprintf("%s/%s/agent-ready", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return

	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}

	defer res.Body.Close()

	v = res.StatusCode == constants.HTTP_STATUS_NO_CONTENT.Integer()
	return
}

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
