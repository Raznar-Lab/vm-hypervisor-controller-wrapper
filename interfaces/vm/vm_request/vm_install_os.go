package vm_request

type VMInstallOSRequestData struct {
	ImageFile string `json:"image_file"`
	MacAddr   string `json:"mac_address"`
}
