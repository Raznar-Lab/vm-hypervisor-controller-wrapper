package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/Raznar-Lab/vm-control-hypervisor-wrapper/interfaces/vm/base_response"
)

type BaseService struct {
	URL         string
	TokenID     string
	TokenSecret string
	Client      *http.Client
}

func (s BaseService) NewHttpRequest(method string, endpoint string, body []byte) (r *http.Request, err error) {
	r, err = http.NewRequest(method, s.URL+path.Clean(fmt.Sprintf("/%s", endpoint)), bytes.NewBuffer(body))
	if err != nil {
		return
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s=%s", s.TokenID, s.TokenSecret))

	return
}

func (s BaseService) HandleErrorResponseNonBody(res *http.Response, expectedStatus int) (err error) {
	defer res.Body.Close()
	if res.StatusCode != expectedStatus {
		var resBody base_response.BaseResponse
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			return fmt.Errorf("unexpected result, expected %d but received %d (%s)", expectedStatus, res.StatusCode, res.Status)
		}
		return fmt.Errorf("unexpected result, expected %d but received %d (%s) - %s", expectedStatus, res.StatusCode, res.Status, resBody.Errors.GetAllString())
	}
	return nil
}

func (s BaseService) HandleErrorResponse(res *base_response.BaseResponse, expectedStatus int) (err error) {
	if res.Status != expectedStatus {
		return fmt.Errorf("unexpected result, expected %d but received %d - %s", expectedStatus, res.Status, res.Errors.GetAllString())
	}

	return nil
}
