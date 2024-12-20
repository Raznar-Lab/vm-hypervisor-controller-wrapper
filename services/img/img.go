package img

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/base"

type ImageService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *ImageService) {
	service = &ImageService{}
	service.URL = bs.URL
	service.TokenSecret = bs.TokenSecret
	service.Client = bs.Client
	return
}
