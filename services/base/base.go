package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/base_response"
)

type BaseService struct {
	URL         string
	TokenSecret string
	Client      *http.Client
}

func (s BaseService) NewHttpRequestJSON(method string, endpoint string, body any) (r *http.Request, err error) {

	resBody, err := json.Marshal(body)
	if err != nil {
		return
	}

	r, err = s.NewHttpRequest(method, endpoint, resBody)
	return
}

func (s BaseService) NewHttpRequest(method string, endpoint string, body []byte) (r *http.Request, err error) {
	r, err = http.NewRequest(method, s.URL+path.Clean(fmt.Sprintf("/%s", endpoint)), bytes.NewBuffer(body))
	if err != nil {
		return
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.TokenSecret))

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
	if res.Code != expectedStatus {
		if res.Errors == nil {
			return fmt.Errorf("unexpected result, expected %d but received %d", expectedStatus, res.Code)
		}
		return fmt.Errorf("unexpected result, expected %d but received %d - %s", expectedStatus, res.Code, res.Errors.GetAllString())
	}

	return nil
}
