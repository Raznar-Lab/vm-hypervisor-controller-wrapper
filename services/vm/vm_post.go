package vm

import (
	"encoding/json"
	"fmt"
	"raznar.id/vm-control-hypervisor-wrapper/interfaces/vm/vm_request"
	"raznar.id/vm-control-hypervisor-wrapper/interfaces/vm/vm_response"
	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s VMService) Create(data vm_request.VMCreateRequestData) (success bool, err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s", constants.ROUTE_VM, data.UUID), body)
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

func (s VMService) Start(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/start", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Restart(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/restart", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Stop(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/stop", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) ForceStop(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/foce-stop", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Suspend(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/suspend", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) Unsuspend(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/unsuspend", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) LinkDisks(uuid string) (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/link-disks", constants.ROUTE_VM, uuid), nil)
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

func (s VMService) ResetPassword(uuid string) (resData *vm_response.VMResetPasswordResponseData, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/reset-password", constants.ROUTE_VM, uuid), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if res.StatusCode != constants.HTTP_STATUS_OK.Integer() {
		if err != nil {
			return
		}

		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_OK, res.StatusCode, res.Status)
		return
	}

	defer res.Body.Close()
	resData = &vm_response.VMResetPasswordResponseData{}
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		resData = nil
		return
	}

	return
}

func (s VMService) SetupNetwork(uuid string, data vm_request.VMSetupNetworkRequestData) (success bool, err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/setup-network", constants.ROUTE_VM, uuid), body)
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

func (s VMService) InstallOS(uuid string, data vm_request.VMInstallOSRequestData) (success bool, err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/install-os", constants.ROUTE_VM, uuid), body)
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

func (s VMService) SwitchBootMode(uuid string, data vm_request.VMBootModeRequestData) (success bool, err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Println(string(body))
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/%s/boot-mode", constants.ROUTE_VM, uuid), body)
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
