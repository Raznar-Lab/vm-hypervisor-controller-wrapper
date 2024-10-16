package img


import "raznar.id/vm-control-hypervisor-wrapper/base"

type ImageService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *ImageService) {
	service = &ImageService{}
	service.URL = bs.URL
	service.TokenID = bs.TokenID
	service.TokenSecret = bs.TokenSecret
	return
}