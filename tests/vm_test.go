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

	var passed uint
	tests := []vmUnitTest{
		{"Create Server", b.createServer},
		{"Update Server", b.updateServer},
		{"Create OS Disk", b.createOsDisk},
		{"Create Second Disk", b.createSecondDisk},
		{"Install OS", b.installOS},
		{"Increase Second Disk Size", b.increaseSecondDiskSize},
		{"Resize OS Disk", b.resizeOsDisk},
		{"Switch Boot to Second", b.switchToOSTOSecond},
		{"Switch Boot to Primary", b.switchToOSBootMode},
		{"Delete Second Disk", b.deleteSecondDisk},
		{"Link Disks", b.linkDisks},
		{"Switch to Recovery Mode", b.switchToRecoveryMode},
		{"Switch to OS Boot Mode", b.switchToOSBootMode},
		{"Get Details", b.getDetails},
		{"Start Server", b.start},
		{"Reset Password", b.resetPassword},
		{"Setup Network", b.setupNetwork},
		{"Restart Server", b.restart},
		{"Suspend Server", b.suspend},
		{"Unsuspend Server", b.unsuspend},
		{"Stop Server", b.stop},
	}

	defer func() {
		if deleteErr := b.deleteServer(vmService, uuidStr); deleteErr != nil {
			b.t.Logf("Error deleting server: %v", deleteErr)
		}

		b.t.Logf("Passed %d of %d server unit tests", passed, len(tests))
	}()

	for _, test := range tests {
		if err = test.Fn(vmService, uuidStr); err != nil {
			b.t.Logf("Error during %s: %v", test.Name, err)
			return err
		}

		passed++
	}

	return nil
}

func (b VMTest) createServer(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Creating server with UUID:", uuidStr)
	success, err := vmService.Create(vm_request.VMCreateData{
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

	time.Sleep(100 * time.Millisecond)
	return nil
}

func (b VMTest) updateServer(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Updating server with UUID:", uuidStr)
	updateData := vm_request.VMUpdateData{
		MacId:       "02:00:00:7d:a8:ab",
		Memory:      4096,
		Cores:       3,
		NetworkRate: 20,
	}

	updateData.SetBaloon(true)
	success, err := vmService.Update(uuidStr, updateData)

	if err != nil {
		b.t.Logf("Error updating server: %v", err)
		return err
	}
	if !success {
		b.t.Log("Server update failed")
		return nil
	}
	b.t.Log("Server updated successfully")

	time.Sleep(100 * time.Millisecond)
	return nil
}

func (b VMTest) installOS(vmService *vm.VMService, uuidStr string) (err error) {
	time.Sleep(100 * time.Millisecond)
	b.t.Log("Installing OS on server:", uuidStr)
	success, err := vmService.InstallOS(uuidStr, vm_request.VMInstallOSRequestData{
		DiskLabel:     "os",
		ImageFile:     "debian-12-amd64.qcow2",
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

func (b VMTest) increaseSecondDiskSize(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Increasing second disk size for server:", uuidStr)
	success, err := vmService.IncreaseDiskSize(uuidStr, vm_request.VMIncreaseDiskSize{
		Label: "second",
		Size:  1,
	})

	if err != nil {
		b.t.Logf("Error increasing second disk size: %v", err)
		return err
	}
	if !success {
		b.t.Log("Second disk size increase failed")
		return nil
	}
	b.t.Log("Second disk size increased successfully")
	return nil
}

func (b VMTest) createOsDisk(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Creating OS disk for server:", uuidStr)
	success, err := vmService.CreateDisk(uuidStr, vm_request.VMCreateDisk{
		Size:          5,
		StorageTarget: "local",
		Label:         "os",
	})

	if err != nil {
		b.t.Logf("Error creating OS disk: %v", err)
		return err
	}
	if !success {
		b.t.Log("OS disk creation failed")
		return nil
	}
	b.t.Log("OS disk created successfully")
	return nil
}

func (b VMTest) resizeOsDisk(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Resizing OS disk for server:", uuidStr)
	success, err := vmService.ResizeDiskSize(uuidStr, vm_request.VMResizeDiskSize{
		Label: "os",
		Size:  10,
	})

	if err != nil {
		b.t.Logf("Error resizing OS disk: %v", err)
		return err
	}
	if !success {
		b.t.Log("OS disk resize failed")
		return nil
	}
	b.t.Log("OS disk resized successfully")
	return nil
}

func (b VMTest) createSecondDisk(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Creating second disk for server:", uuidStr)
	success, err := vmService.CreateDisk(uuidStr, vm_request.VMCreateDisk{
		Size:          5,
		StorageTarget: "local",
		Label:         "second",
	})

	if err != nil {
		b.t.Logf("Error creating second disk: %v", err)
		return err
	}
	if !success {
		b.t.Log("Second disk creation failed")
		return nil
	}
	b.t.Log("Second disk created successfully")
	return nil
}

func (b VMTest) deleteSecondDisk(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Deleting second disk for server:", uuidStr)
	success, err := vmService.DeleteDisk(uuidStr, vm_request.VMDeleteDisk{
		Label: "second",
	})

	if err != nil {
		b.t.Logf("Error deleting second disk: %v", err)
		return err
	}
	if !success {
		b.t.Log("Second disk deletion failed")
		return nil
	}
	b.t.Log("Second disk deleted successfully")
	return nil
}

func (b VMTest) linkDisks(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Linking disks for server:", uuidStr)
	success, err := vmService.LinkDisks(uuidStr)

	if err != nil {
		b.t.Logf("Error linking disks: %v", err)
		return err
	}
	if !success {
		b.t.Log("Disk linking failed")
		return nil
	}
	b.t.Log("Disks linked successfully")
	return nil
}

func (b VMTest) switchToRecoveryMode(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Switching boot mode to recovery on server:", uuidStr)
	success, err := vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootDiskLabel:       "os",
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
		BootDiskLabel: "os",
		BootMode:      "os",
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

func (b VMTest) switchToOSTOSecond(vmService *vm.VMService, uuidStr string) (err error) {
	b.t.Log("Switching second disk to primary:", uuidStr)
	success, err := vmService.SwitchBootMode(uuidStr, vm_request.VMBootModeRequestData{
		BootDiskLabel: "second",
		BootMode:      "os",
	})

	if err != nil {
		b.t.Logf("Error switching boot mode to second disk: %v", err)
		return err
	}
	if !success {
		b.t.Log("Switch to second disk failed")
		return nil
	}
	b.t.Log("Switched to second disk successfully")
	time.Sleep(2500 * time.Millisecond)
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
	time.Sleep(2 * time.Second)
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
