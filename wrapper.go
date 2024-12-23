package hypervisor_controller_wrapper

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/wrapper"
)

func New(url string, tokenSecret string) (sWrapper *wrapper.Wrapper) {
	sWrapper = wrapper.New(url, tokenSecret)
	return
}
