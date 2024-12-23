package nw

import (
	"fmt"
	"strings"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/nw/nw_request"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s NetworkService) Delete(ipv4 string, macId string) (success bool, err error) {
	baseURL := fmt.Sprintf("%s/delete", constants.ROUTE_NW)

	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, baseURL, nw_request.NWDeleteIPNetworkRequest{
		IPv4Cidr: ipv4,
		MacID:    macId,
	})

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

func (s NetworkService) Add(ipv4cidr string, macid string) (success bool, err error) {
	baseURL := constants.ROUTE_NW
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, baseURL, nw_request.NWAddIPNetworkRequest{
		IPv4Cidr: ipv4cidr,
		MacID:    macid,
	})
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

func (s NetworkService) CreateMultiple(ipv4List []string, macId string) (success bool, err error) {
	baseURL := fmt.Sprintf("%s/multiple", constants.ROUTE_NW)

	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, baseURL, nw_request.NWAddMultipleIPNetworkRequest{
		MacID:        macId,
		IPv4CidrList: strings.Join(ipv4List, ","),
	})
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
