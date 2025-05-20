package vm_request

type VMBootModeRequestData struct {
	BootDiskLabel        string `json:"boot_disk_label"`
	BootMode             string `json:"boot_mode"`
	BootRecoveryImageURL string `json:"boot_recovery_image_url"`
	BootRecoveryStorage  string `json:"boot_recovery_storage"`
}
