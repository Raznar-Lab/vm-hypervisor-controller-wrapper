package hypervisor_controller_wrapper

import (
	"os"
	"testing"
	"time"

	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/v1/interfaces/vm/vm_request"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type VMTest struct {
	BaseTest
}

func (b VMTest) Start() (err error) {
	uuidStr := uuid.NewString()

	vmService := b.wrapper.VM()
	b.t.Log("Creating the server")
	success, err := vmService.Create(vm_request.VMCreateRequestData{
		UUID:    uuidStr,
		MacId:   "02:00:00:7d:a8:ad",
		Memory:  2048,
		Balloon: false,
		Cores:   2,
	})

	if err != nil {
		b.t.Logf("Failed to create the server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server creation was not successful")
		return nil
	}

	b.t.Log("Successfully created the server")

	time.Sleep(100 * time.Millisecond)
	// Installing OS
	b.t.Log("Installing OS on the server")
	success, err = vmService.InstallOS(uuidStr, vm_request.VMInstallOSRequestData{
		OSFile:        "rockylinux-9-amd64.qcow2",
		StorageTarget: "local",
		DiskSize:      20,
	})

	if err != nil {
		b.t.Logf("Failed to install OS: %v", err)
		return err
	}
	if !success {
		b.t.Log("OS installation was not successful")
		return nil
	}

	b.t.Log("Successfully installed OS on the server")

	// Switching boot mode
	b.t.Log("Switching boot mode to recovery")
	success, err = vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootMode:            "recovery",
		BootRecoveryOS:      "ubuntu-24-04-amd64.qcow2",
		BootRecoveryStorage: "local",
	})

	if err != nil {
		b.t.Logf("Failed to switch boot mode: %v", err)
		return err
	}
	if !success {
		b.t.Log("Boot mode switching was not successful")
		return nil
	}

	b.t.Log("Successfully switched to recovery mode")

	// Switching boot mode to OS
	b.t.Log("Switching boot mode to os")
	success, err = vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootMode: "os",
	})

	if err != nil {
		b.t.Logf("Failed to switch boot mode: %v", err)
		return err
	}
	if !success {
		b.t.Log("Boot mode switching was not successful")
		return nil
	}

	b.t.Log("Successfully switched to os mode")

	// Deleting the VM
	b.t.Log("Deleting the server")
	success, err = vmService.Delete(uuidStr)

	if err != nil {
		b.t.Logf("Failed to delete the server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server deletion was not successful")
		return nil
	}

	b.t.Log("Successfully deleted the server")
	return nil
}

func TestVM(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	apiURL := os.Getenv("API_URL")
	apiTokenID := os.Getenv("API_TOKEN_ID")
	apiToken := os.Getenv("API_TOKEN")
	wrapper := New(apiURL, apiTokenID, apiToken)
	vmTest := VMTest{}
	vmTest.wrapper = wrapper
	vmTest.t = t

	err = vmTest.Start()
	if err != nil {
		panic(err)
	}
}
