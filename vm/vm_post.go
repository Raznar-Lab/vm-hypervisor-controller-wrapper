package vm

type VMCreateRequestData struct {
	UUID    string `json:"uuid"`
	MacId   string `json:"mac_id"`
	Memory  int    `json:"memory"`
	Cores   int    `json:"cores"`
	Balloon bool   `json:"balloon"`
}

func (s VMService) Create(data VMCreateRequestData) (success bool) {
	return
}