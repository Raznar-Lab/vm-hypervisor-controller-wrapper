package main

import (
	"fmt"
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

func main() {
	wrapper := New("http://localhost:8080", "cikoHNJOA2", "Caiktjhoi@tha@ithai@tha@izz")
	data, err := wrapper.IMG().Download("https://gist.githubusercontent.com/NotYusta/007c3f579c3d062775367f910fb8b29e/raw/d303a349f25044d91dd09933bb7b88a05a839cc7/reset_ssh_password.sh", "download.ssh")
	if err != nil {
		return
	}

	wrapper.IMG().Delete("download.ssh")

	fmt.Println(data)
}