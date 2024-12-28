package vm_request

type VMInstallOSRequestData struct {
	DiskLabel     string `json:"disk_label"`
	ImageFile     string `json:"image_file"`
}
