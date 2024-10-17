package vm_request

type VMInstallOSRequestData struct {
	OSFile        string `json:"os_file"`
	StorageTarget string `json:"storage_target"`
	DiskSize      int64  `json:"disk_size"`
}
