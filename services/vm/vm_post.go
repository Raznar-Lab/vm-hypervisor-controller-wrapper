package vm

import (
	"encoding/json"
	"fmt"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/vm_request"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/vm_response"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s VMService) Create(data vm_request.VMCreateRequestData) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, constants.ROUTE_VM, data)
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

func (s VMService) IncreaseDiskSize(uuid string, data vm_request.VMIncreaseDiskSize) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/disk/increase", constants.ROUTE_VM, uuid), data)
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

func (s VMService) ResizeDiskSize(uuid string, data vm_request.VMResizeDiskSize) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/disk/resize", constants.ROUTE_VM, uuid), data)
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

func (s VMService) CreateDisk(uuid string, data vm_request.VMCreateDisk) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/disk", constants.ROUTE_VM, uuid), data)
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

func (s VMService) DeleteDisk(uuid string, data vm_request.VMDeleteDisk) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/disk/delete", constants.ROUTE_VM, uuid), data)
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

func (s VMService) Start(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/start", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Restart(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/restart", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Stop(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/stop", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) ForceStop(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/force-stop", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Suspend(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/suspend", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Unsuspend(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/unsuspend", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) LinkDisks(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/disk/link", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) ResetPassword(uuid string, data vm_request.VMResetPasswordRequestData) (resData *vm_response.VMResetPasswordResponseData, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/reset-password", constants.ROUTE_VM, uuid), data)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}

	resData = &vm_response.VMResetPasswordResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	err = s.HandleErrorResponse(&resData.BaseResponse, constants.HTTP_STATUS_OK.Integer())
	return
}

func (s VMService) SetupNetwork(uuid string, data vm_request.VMSetupNetworkRequestData) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/setup-network", constants.ROUTE_VM, uuid), data)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_OK.Integer())
	if err != nil {
		return false, err
	}

	success = true
	return
}

func (s VMService) InstallOS(uuid string, data vm_request.VMInstallOSRequestData) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/install-os", constants.ROUTE_VM, uuid), data)
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

func (s VMService) SwitchBootMode(uuid string, data vm_request.VMBootModeRequestData) (success bool, err error) {
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/boot-mode", constants.ROUTE_VM, uuid), data)
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
