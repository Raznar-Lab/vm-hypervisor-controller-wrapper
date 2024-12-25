package vm_request

type VMIncreaseDiskSize struct {
	Label string `json:"label"`
	Size  int    `json:"size"`
}

type VMResizeDiskSize struct {
	Label string `json:"label"`
	Size  int    `json:"size"`
}

type VMCreateDisk struct {
	Label         string `json:"label"`
	StorageTarget string `json:"storage_target"`
	Size          int    `json:"size"`
}

type VMDeleteDisk struct {
	Label string `json:"label"`
}
