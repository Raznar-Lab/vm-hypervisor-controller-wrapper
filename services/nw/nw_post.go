package nw

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/dev/pkg/constants"
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
	defer res.Body.Close()

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_NO_CONTENT.Integer())
	if err != nil {
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
	defer res.Body.Close()

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_NO_CONTENT.Integer())
	if err != nil {
		return
	}

	success = true
	return
}

func (s NetworkService) CreateMultiple(ipv4List []string, macid string) (success bool, err error) {
	baseURL := fmt.Sprintf("%s/multiple", constants.ROUTE_NW)
	params := url.Values{}
	params.Add("ipv4-list", strings.Join(ipv4List, ","))
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
	defer res.Body.Close()

	err = s.HandleErrorResponseNonBody(res, constants.HTTP_STATUS_NO_CONTENT.Integer())
	if err != nil {
		return
	}

	success = true
	return
}
