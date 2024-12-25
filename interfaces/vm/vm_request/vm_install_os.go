package vm_request

type VMInstallOSRequestData struct {
	Label         string `json:"label"`
	ImageFile     string `json:"image_file"`
	StorageTarget string `json:"storage_target"`
	DiskSize      int    `json:"disk_size"`
}
