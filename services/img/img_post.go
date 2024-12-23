package img

import (
	"encoding/json"
	"fmt"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/img/img_request"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/img/img_response"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/pkg/constants"
)

func (s ImageService) GetAll() (resData img_response.IMGGetAllResponseData, err error) {
	// Build the URL using net/url to properly encode parameters
	baseURL := constants.ROUTE_IMG
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_GET, baseURL, nil)
	if err != nil {
		return
	}

	res, err := s.Client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resData)
	if err != nil {
		return
	}

	err = s.HandleErrorResponse(&resData.BaseResponse, constants.HTTP_STATUS_OK.Integer())
	return
}

func (s ImageService) Download(downloadURL string, filename string) (success bool, err error) {
	// Build the URL using net/url to properly encode parameters
	baseURL := fmt.Sprintf("%s/download", constants.ROUTE_IMG)

	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, baseURL, img_request.IMGDownloadImageRequest{
		DownloadURL: downloadURL,
		Filename:    filename,
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

func (s ImageService) Delete(filename string) (success bool, err error) {
	baseURL := fmt.Sprintf("%s/delete", constants.ROUTE_IMG)
	r, err := s.NewHttpRequestJSON(constants.HTTP_METHOD_POST, baseURL, img_request.IMGDeleteImageRequest{
		Filename: filename,
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

func (s ImageService) Clear() (success bool, err error) {
	baseURL := fmt.Sprintf("%s/clear", constants.ROUTE_IMG)
	r, err := s.NewHttpRequest(constants.HTTP_METHOD_POST, baseURL, nil)

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
