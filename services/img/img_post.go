package img

import (
	"fmt"
	"net/url"
	"raznar.id/vm-control-hypervisor-wrapper/pkg/constants"
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
	defer res.Body.Close() // Ensure the response body is closed

	if res.StatusCode != constants.HTTP_STATUS_NO_CONTENT.Integer() {
		err = fmt.Errorf("unexpected result, expected %d but received %d (%s)", constants.HTTP_STATUS_NO_CONTENT, res.StatusCode, res.Status)
		return
	}

	success = true
	return
}
