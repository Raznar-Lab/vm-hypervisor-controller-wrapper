package img

import (
	"fmt"
	"net/url"
	"github.com/Raznar-Lab/vm-control-hypervisor-wrapper/pkg/constants"
)

func (s ImageService) Download(downloadURL string, filename string) (success bool, err error) {
	// Build the URL using net/url to properly encode parameters
	baseURL := fmt.Sprintf("%s/download", constants.ROUTE_IMG)
	params := url.Values{}
	params.Add("download_url", downloadURL)
	params.Add("filename", filename)

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
