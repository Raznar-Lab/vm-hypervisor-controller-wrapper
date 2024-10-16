package vm_response

type VMMetricsResponseData struct {
	Status int           `json:"status"`
	Data   vmMetricsData `json:"data"`
}

type vmMetricsData struct {
	NetIn  float64 `json:"net_in"`  // Received bytes
	NetOut float64 `json:"net_out"` // Sent bytes
}
