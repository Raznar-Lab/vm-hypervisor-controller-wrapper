package nw

import (
	"fmt"
	"net/url"
	"strings"

	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s NetworkService) Reset() (success bool, err error) {
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fmt.Sprintf("%s/reset", constants.ROUTE_NW), nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close() // Ensure the response body is closed

	if res.StatusCode != constants.HTTP_STATUS_NO_CONTENT.Integer() {
		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_NO_CONTENT, res.StatusCode, res.Status)
		return
	}

	success = true
	return
}

func (s NetworkService) Create(ipv4 string, macid string) (success bool, err error) {
	baseURL := constants.ROUTE_NW
	params := url.Values{}
	params.Add("ipv4", ipv4)
	params.Add("macid", macid)

	// Create the full request URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fullURL, nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close() // Ensure the response body is closed

	if res.StatusCode != constants.HTTP_STATUS_NO_CONTENT.Integer() {
		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_NO_CONTENT, res.StatusCode, res.Status)
		return
	}

	success = true
	return
}

func (s NetworkService) CreateMultiple(ipv4List []string, macid string) (success bool, err error) {
	baseURL := constants.ROUTE_NW
	params := url.Values{}
	params.Add("ipv4_list", strings.Join(ipv4List, ","))
	params.Add("macid", macid)

	// Create the full request URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, fullURL, nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close() // Ensure the response body is closed

	if res.StatusCode != constants.HTTP_STATUS_NO_CONTENT.Integer() {
		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_NO_CONTENT, res.StatusCode, res.Status)
		return
	}

	success = true
	return
}
