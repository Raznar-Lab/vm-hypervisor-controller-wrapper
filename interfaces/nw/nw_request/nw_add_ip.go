package nw_request

type NWAddIPNetworkRequest struct {
	IPv4Cidr string `json:"ipv4_cidr"`
	MacID    string `json:"mac_id"`
}

type NWAddMultipleIPNetworkRequest struct {
	IPv4CidrList string `json:"ipv4_cidr_list"`
	MacID        string `json:"mac_id"`
}
