package wrapper

import (
	"net/http"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/base"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/img"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/nw"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/vm"
)

type Wrapper struct {
	base.BaseService
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

func New(url string, tokenSecret string) (wrapper *Wrapper) {
	wrapper = &Wrapper{}
	wrapper.URL = url
	wrapper.TokenSecret = tokenSecret
	wrapper.Client = http.DefaultClient
	return
}
