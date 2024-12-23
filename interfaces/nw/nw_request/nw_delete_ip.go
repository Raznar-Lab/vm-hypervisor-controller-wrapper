package nw_request

type NWDeleteIPNetworkRequest struct {
	IPv4Cidr string `json:"ipv4_cidr"`
	MacID    string `json:"mac_id"`
}
