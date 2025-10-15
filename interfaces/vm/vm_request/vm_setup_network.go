package vm_request

type VMSetupNetworkRequestData struct {
	OSType  string `json:"os_type"`
	IP      string `json:"ip"`
	IpCIDR  string `json:"ip_cidr"`
	Gateway string `json:"gateway"`
	DNS1    string `json:"dns1"`
	DNS2    string `json:"dns2"`

	IPv6         string `json:"ip_v6"`
	IpCIDRV6     string `json:"ip_cidr_v6"`
	GatewayV6    string `json:"gateway_v6"`
	DNS1V6       string `json:"dns1_v6"`
	DNS2V6       string `json:"dns2_v6"`
	MacAddress string `json:"mac_address"`
}
