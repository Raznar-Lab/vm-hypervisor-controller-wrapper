package vm

import (
	"net/http"
	"raznar.id/vm-control-hypervisor-wrapper/services/base"
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