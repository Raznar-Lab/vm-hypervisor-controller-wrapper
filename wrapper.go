package hypervisor_controller_wrapper

import (
	"net/http"
	"raznar.id/vm-control-hypervisor-wrapper/services/base"
	"raznar.id/vm-control-hypervisor-wrapper/services/img"
	"raznar.id/vm-control-hypervisor-wrapper/services/nw"
	"raznar.id/vm-control-hypervisor-wrapper/services/vm"
)

type Wrapper struct {
	base.BaseService
}

func New(url string, tokenId string, tokenSecret string) (wrapper *Wrapper) {
	wrapper = &Wrapper{}
	wrapper.URL = url
	wrapper.TokenID = tokenId
	wrapper.TokenSecret = tokenSecret
	wrapper.Client = http.DefaultClient
	return
}

func (w Wrapper) IMG() (service *img.ImageService) {
	service = img.New(&w.BaseService)
	return
}

func (w Wrapper) NW() (service *nw.NetworkService) {
	service = nw.New(&w.BaseService)
	return
}

func (w Wrapper) VM() (service *vm.VMService) {
	service = vm.New(&w.BaseService)
	return
}