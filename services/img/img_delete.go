package img

import (
	"fmt"
	"net/url"

	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s ImageService) Delete(filename string) (success bool, err error) {
	baseURL := constants.ROUTE_IMG
	params := url.Values{}
	params.Add("filename", filename)

	// Create the full request URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	r, err := s.NewHttpRequest(constants.HTTP_METHOD_DELETE, fullURL, nil)
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
