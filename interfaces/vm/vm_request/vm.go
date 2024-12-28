package vm_request

type VMCreateData struct {
	UUID        string `json:"uuid"`
	MacId       string `json:"mac_id"`
	Memory      int    `json:"memory"`
	Cores       int    `json:"cores"`
	Balloon     bool   `json:"balloon"`
	NetworkRate int64  `json:"network_rate"`
}

type VMUpdateData struct {
	MacId       string `json:"mac_id"`
	Memory      int64  `json:"memory"`
	Cores       int64  `json:"cores"`
	Balloon     *bool  `json:"balloon"`
	NetworkRate int64  `json:"network_rate"`
}

func (d VMUpdateData) SetBaloon(value bool) {
	d.Balloon = &value
}
