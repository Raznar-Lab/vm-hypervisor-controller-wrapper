package vm

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/dev/services/base"
	"net/http"
)

type VMService struct {
	base.BaseService
}

func New(bs *base.BaseService) (service *VMService) {
	service = &VMService{}
	service.URL = bs.URL
	service.TokenID = bs.TokenID
	service.TokenSecret = bs.TokenSecret
	service.Client = &http.Client{}
	return
}
