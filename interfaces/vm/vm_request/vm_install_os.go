package vm_request

type VMInstallOSRequestData struct {
	ImageURL string `json:"image_url"`
	MacAddr  string `json:"mac_address"`
}
