package vm_request


type VMSetupNetworkRequestData struct {
	OSType     string `json:"os_type"`
	IPv4       string `json:"ipv4"`
	IpCIDR     string `json:"ip_cidr"`
	Gateway    string `json:"gateway"`
	DNS1       string `json:"dns1"`
	DNS2       string `json:"dns2"`
	MacAddress string `json:"mac_address"`
}
