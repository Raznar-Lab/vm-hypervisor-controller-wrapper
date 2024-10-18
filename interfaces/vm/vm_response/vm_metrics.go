package vm_response

import "raznar.id/vm-control-hypervisor-wrapper/interfaces/vm/base_response"

type VMMetricsResponseData struct {
	base_response.BaseResponse
	Data vmMetricsData `json:"data"`
}

type vmMetricsData struct {
	NetIn  float64 `json:"net_in"`  // Received bytes
	NetOut float64 `json:"net_out"` // Sent bytes
}
