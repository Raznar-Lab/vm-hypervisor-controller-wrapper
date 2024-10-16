package vm_request

type VMBootModeRequestData struct {
	BootMode            string `json:"boot_mode"`
	BootRecoveryOS      string `json:"boot_recovery_os"`
	BootRecoveryStorage string `json:"boot_recovery_storage"`
}
