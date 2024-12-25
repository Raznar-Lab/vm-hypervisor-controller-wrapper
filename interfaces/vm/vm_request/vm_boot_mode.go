package vm_request

type VMBootModeRequestData struct {
	BootDiskLabel       string `json:"boot_disk_label"`
	BootMode            string `json:"boot_mode"`
	BootRecoveryImage   string `json:"boot_recovery_image"`
	BootRecoveryStorage string `json:"boot_recovery_storage"`
}
