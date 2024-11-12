package img

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/base"

type ImageService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *ImageService) {
	service = &ImageService{}
	service.URL = bs.URL
	service.TokenID = bs.TokenID
	service.TokenSecret = bs.TokenSecret
	service.Client = bs.Client
	return
}
