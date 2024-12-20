package vm

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/base"
	"net/http"
)

type VMService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *VMService) {
	service = &VMService{}
	service.URL = bs.URL
	service.TokenSecret = bs.TokenSecret
	service.Client = &http.Client{}
	return
}
