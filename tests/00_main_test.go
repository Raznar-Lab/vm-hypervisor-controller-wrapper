package tests

import (
	"os"
	"path"
	"testing"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/wrapper"
	"github.com/joho/godotenv"
)

type BaseTest struct {
	wrapper *wrapper.Wrapper
	t       *testing.T
}

func loadEnv() (sWrapper *wrapper.Wrapper, err error) {

	err = godotenv.Load(path.Join("../.env"))
	if err != nil {
		return
	}

	apiURL := os.Getenv("API_URL")
	apiToken := os.Getenv("API_TOKEN")
	sWrapper = wrapper.New(apiURL, apiToken)
	return
}
