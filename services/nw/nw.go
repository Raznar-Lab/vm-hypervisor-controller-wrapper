package nw

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/base"
)

type NetworkService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *NetworkService) {
	service = &NetworkService{}
	service.URL = bs.URL
	service.TokenSecret = bs.TokenSecret
	service.Client = bs.Client
	return
}
