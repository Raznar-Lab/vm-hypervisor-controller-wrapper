package nw

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/v1/services/base"
)

type NetworkService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *NetworkService) {
	service = &NetworkService{}
	service.URL = bs.URL
	service.TokenID = bs.TokenID
	service.TokenSecret = bs.TokenSecret
	service.Client = bs.Client
	return
}
