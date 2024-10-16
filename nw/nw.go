package nw


import "raznar.id/vm-control-hypervisor-wrapper/base"

type NetworkService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *NetworkService) {
	service = &NetworkService{}
	service.URL = bs.URL
	service.TokenID = bs.TokenID
	service.TokenSecret = bs.TokenSecret
	return
}