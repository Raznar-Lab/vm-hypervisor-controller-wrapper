package vm_request

type VMInstallOSRequestData struct {
	OSFile        string `json:"os_file"`
	StorageTarget string `json:"storage_target"`
	DiskSize      string `json:"disk_size"`
}
