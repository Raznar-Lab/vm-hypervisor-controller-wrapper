package tests

import (
	"testing"
	"time"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/vm_request"
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/vm"
	"github.com/google/uuid"
)

type VMTest struct {
	BaseTest
}

type vmUnitTest struct {
	Name string
	Fn   func(*vm.VMService, string) error
}

func (b VMTest) Start() (err error) {
	uuidStr := uuid.NewString()
	vmService := b.wrapper.VM()

	defer func() {
		if deleteErr := b.deleteServer(vmService, uuidStr); deleteErr != nil {
			b.t.Logf("Error deleting server: %v", deleteErr)
		}
	}()

	tests := []vmUnitTest{
		{"Create Server", b.createServer},
		{"Install OS", b.installOS},
		{"Increase Disk Size", b.increaseDiskSize},
		{"Switch to Recovery Mode", b.switchToRecoveryMode},
		{"Switch to OS Boot Mode", b.switchToOSBootMode},
		{"Get Details", b.getDetails},
		{"Start Server", b.start},
		{"Reset Password", b.resetPassword},
		{"Setup Network", b.setupNetwork},
		{"Restart Server", b.restart},
		{"Suspend Server", b.suspend},
		{"Suspend Server", b.unsuspend},
		{"Stop Server", b.stop},
	}

	for _, test := range tests {
		if err = test.Fn(vmService, uuidStr); err != nil {
			b.t.Logf("Error during %s: %v", test.Name, err)
			return err
		}
	}

	return nil
}

func (b VMTest) createServer(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Creating server with UUID:", uuidStr)
	success, err := vmService.Create(vm_request.VMCreateRequestData{
		UUID:    uuidStr,
		MacId:   "02:00:00:7d:a8:ad",
		Memory:  2048,
		Balloon: false,
		Cores:   2,
	})

	if err != nil {
		b.t.Logf("Error creating server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server creation failed")
		return nil
	}
	b.t.Log("Server created successfully")
	return nil
}

func (b VMTest) installOS(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(100 * time.Millisecond)
	b.t.Log("Installing OS on server:", uuidStr)
	success, err := vmService.InstallOS(uuidStr, vm_request.VMInstallOSRequestData{
		OSFile:        "debian-12-amd64.qcow2",
		StorageTarget: "local",
		DiskSize:      5,
	})

	if err != nil {
		b.t.Logf("Error installing OS: %v", err)
		return err
	}
	if !success {
		b.t.Log("OS installation failed")
		return nil
	}
	b.t.Log("OS installed successfully")
	return nil
}

func (b VMTest) increaseDiskSize(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Increasing disk size for server:", uuidStr)
	success, err := vmService.IncreaseDiskSize(uuidStr, vm_request.VMIncreaseDiskSize{
		Size: 1,
	})

	if err != nil {
		b.t.Logf("Error increasing disk size: %v", err)
		return err
	}
	if !success {
		b.t.Log("Disk size increase failed")
		return nil
	}
	b.t.Log("Disk size increased successfully")
	return nil
}

func (b VMTest) switchToRecoveryMode(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Switching boot mode to recovery on server:", uuidStr)
	success, err := vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootMode:            "recovery",
		BootRecoveryImage:   "debian-12-amd64.qcow2",
		BootRecoveryStorage: "local",
	})

	if err != nil {
		b.t.Logf("Error switching boot mode to recovery: %v", err)
		return err
	}
	if !success {
		b.t.Log("Switch to recovery mode failed")
		return nil
	}
	b.t.Log("Switched to recovery mode successfully")
	return nil
}

func (b VMTest) switchToOSBootMode(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Switching boot mode back to OS on server:", uuidStr)
	success, err := vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootMode: "os",
	})

	if err != nil {
		b.t.Logf("Error switching boot mode to OS: %v", err)
		return err
	}
	if !success {
		b.t.Log("Switch to OS mode failed")
		return nil
	}
	b.t.Log("Switched to OS mode successfully")
	return nil
}

func (b VMTest) getDetails(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(200 * time.Millisecond)
	b.t.Log("Getting details for server:", uuidStr)
	details, err := vmService.GetDetails(uuidStr)

	if err != nil {
		b.t.Logf("Error getting server details: %v", err)
		return err
	}
	b.t.Logf("Server details: %v", details)
	return nil
}

func (b VMTest) resetPassword(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(200 * time.Millisecond)
	b.t.Log("Resetting password for server:", uuidStr)
	resData, err := vmService.ResetPassword(uuidStr, vm_request.VMResetPasswordRequestData{
		OSType: "linux",
	})

	if err != nil {
		b.t.Logf("Error resetting password: %v", err)
		return err
	}

	b.t.Logf("Password reset successfully: (u: %s, p: %s)", resData.Data.Username, resData.Data.Password)
	return nil
}

func (b VMTest) setupNetwork(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Setting up network for server:", uuidStr)

	time.Sleep(200 * time.Millisecond) // Add a delay to ensure the VM is ready for network setup.

	success, err := vmService.SetupNetwork(uuidStr, vm_request.VMSetupNetworkRequestData{
		OSType:     "linux",             // Example OS type; update as needed
		IPv4:       "192.168.1.100",     // Example IP address
		IpCIDR:     "24",                // Example subnet CIDR
		Gateway:    "192.168.1.1",       // Example gateway
		DNS1:       "1.1.1.1",           // Primary DNS
		DNS2:       "8.8.8.8",           // Secondary DNS
		MacAddress: "02:00:00:7d:a8:ad", // Example MAC address
	})

	if err != nil {
		b.t.Logf("Error setting up network: %v", err)
		return err
	}
	if !success {
		b.t.Log("Network setup failed")
		return nil
	}
	b.t.Log("Network setup completed successfully")
	return nil
}

func (b VMTest) start(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(500 * time.Millisecond)
	b.t.Log("Starting server:", uuidStr)
	success, err := vmService.Start(uuidStr)

	if err != nil {
		b.t.Logf("Error starting server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server start failed")
		return nil
	}

	time.Sleep(15 * time.Second)
	b.t.Log("Server started successfully")

	return nil
}

func (b VMTest) restart(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(500 * time.Millisecond)
	b.t.Log("Restarting server:", uuidStr)
	success, err := vmService.Restart(uuidStr)

	if err != nil {
		b.t.Logf("Error restarting server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server restart failed")
		return nil
	}
	b.t.Log("Server restarted successfully")
	return nil
}

func (b VMTest) suspend(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(500 * time.Millisecond)
	b.t.Log("Suspending server:", uuidStr)
	success, err := vmService.Suspend(uuidStr)

	if err != nil {
		b.t.Logf("Error suspending server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server suspend failed")
		return nil
	}
	b.t.Log("Server suspended successfully")
	return nil
}

func (b VMTest) unsuspend(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(500 * time.Millisecond)
	b.t.Log("Unsuspending server:", uuidStr)
	success, err := vmService.Unsuspend(uuidStr)

	if err != nil {
		b.t.Logf("Error unsuspending server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server unsuspend failed")
		return nil
	}
	b.t.Log("Server unsuspended successfully")
	return nil
}

func (b VMTest) stop(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(500 * time.Millisecond)
	b.t.Log("Stopping server:", uuidStr)
	success, err := vmService.Stop(uuidStr)

	if err != nil {
		b.t.Logf("Error stopping server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server stop failed")
		return nil
	}
	b.t.Log("Server stopped successfully, finale.")
	time.Sleep(5 * time.Second)
	return nil
}

func (b VMTest) deleteServer(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Deleting server:", uuidStr)
	success, err := vmService.Delete(uuidStr)

	if err != nil {
		b.t.Logf("Error deleting server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server deletion failed")
		return nil
	}
	b.t.Log("Server deleted successfully")
	return nil
}

func TestVM(t *testing.T) {
	// Load environment variables
	wrapper, err := loadEnv()
	if err != nil {
		t.Fatalf("cannot load wrapper!: %s", err.Error())
		return
	}
	
	vmTest := VMTest{}
	vmTest.wrapper = wrapper
	vmTest.t = t

	// Start the test
	err = vmTest.Start()
	if err != nil {
		panic(err)
	}
}
