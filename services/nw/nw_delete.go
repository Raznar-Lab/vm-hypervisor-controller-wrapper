package nw

import (
	"fmt"
	"net/url"

	"github.com/Raznar-Lab/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s NetworkService) Delete(ipv4 string, macid string) (success bool, err error) {
	baseURL := constants.ROUTE_NW
	params := url.Values{}
	params.Add("ipv4", ipv4)
	params.Add("macid", macid)


	// Create the full request URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_DELETE, fullURL, nil)
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
