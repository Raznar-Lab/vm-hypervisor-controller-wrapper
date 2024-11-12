package vm_response

import "github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/interfaces/vm/base_response"

type VMDetailsResponseData struct {
	base_response.BaseResponse
	Data vmDetailsData `json:"data"`
}

type vmDetailsData struct {
	Mem            int64                       `json:"mem"`
	Disk           int64                       `json:"disk"`
	CPUs           int                         `json:"cpus"`
	RunningQEMU    string                      `json:"running-qemu"`
	VMID           int                         `json:"vmid"`
	MaxMem         int64                       `json:"maxmem"`
	NetIn          int64                       `json:"netin"`
	Uptime         int64                       `json:"uptime"`
	CPU            float64                     `json:"cpu"`
	DiskWrite      int64                       `json:"diskwrite"`
	QMPStatus      string                      `json:"qmpstatus"`
	Nics           map[string]detailsNIC       `json:"nics"`
	MaxDisk        int64                       `json:"maxdisk"`
	DiskRead       int64                       `json:"diskread"`
	HA             detailsHAInfo               `json:"ha"`
	ProxmoxSupport detailsProxmoxSupport       `json:"proxmox-support"`
	Name           string                      `json:"name"`
	RunningMachine string                      `json:"running-machine"`
	BlockStat      map[string]detailsBlockStat `json:"blockstat"`
	PID            int                         `json:"pid"`
	Agent          int                         `json:"agent"`
	NetOut         int64                       `json:"netout"`
	Status         string                      `json:"status"`
}

type detailsNIC struct {
	NetOut int64 `json:"netout"`
	NetIn  int64 `json:"netin"`
}

type detailsHAInfo struct {
	Managed int `json:"managed"`
}

type detailsProxmoxSupport struct {
	PBSDirtyBitmapSaveVM    bool   `json:"pbs-dirty-bitmap-savevm"`
	PBSDirtyBitmap          bool   `json:"pbs-dirty-bitmap"`
	BackupFleecing          bool   `json:"backup-fleecing"`
	PBSLibraryVersion       string `json:"pbs-library-version"`
	PBSDirtyBitmapMigration bool   `json:"pbs-dirty-bitmap-migration"`
	PBSMasterKey            bool   `json:"pbs-masterkey"`
	BackupMaxWorkers        bool   `json:"backup-max-workers"`
	QueryBitmapInfo         bool   `json:"query-bitmap-info"`
}

type detailsBlockStat struct {
	RdMerged                    int   `json:"rd_merged"`
	InvalidFlushOperations      int   `json:"invalid_flush_operations"`
	AccountInvalid              bool  `json:"account_invalid"`
	WrMerged                    int   `json:"wr_merged"`
	FailedZoneAppendOperations  int   `json:"failed_zone_append_operations"`
	InvalidUnmapOperations      int   `json:"invalid_unmap_operations"`
	FailedFlushOperations       int   `json:"failed_flush_operations"`
	TimedStats                  []int `json:"timed_stats"`
	ZoneAppendOperations        int   `json:"zone_append_operations"`
	IdleTimeNs                  int64 `json:"idle_time_ns"`
	UnmapMerged                 int   `json:"unmap_merged"`
	FailedRdOperations          int   `json:"failed_rd_operations"`
	WrHighestOffset             int64 `json:"wr_highest_offset"`
	InvalidZoneAppendOperations int   `json:"invalid_zone_append_operations"`
	FailedUnmapOperations       int   `json:"failed_unmap_operations"`
	AccountFailed               bool  `json:"account_failed"`
	UnmapTotalTimeNs            int64 `json:"unmap_total_time_ns"`
	WrTotalTimeNs               int64 `json:"wr_total_time_ns"`
	RdTotalTimeNs               int64 `json:"rd_total_time_ns"`
	FlushTotalTimeNs            int64 `json:"flush_total_time_ns"`
	RdBytes                     int64 `json:"rd_bytes"`
	WrBytes                     int64 `json:"wr_bytes"`
	InvalidRdOperations         int   `json:"invalid_rd_operations"`
	ZoneAppendBytes             int64 `json:"zone_append_bytes"`
	ZoneAppendMerged            int   `json:"zone_append_merged"`
	FailedWrOperations          int   `json:"failed_wr_operations"`
	ZoneAppendTotalTimeNs       int64 `json:"zone_append_total_time_ns"`
	FlushOperations             int   `json:"flush_operations"`
	UnmapOperations             int   `json:"unmap_operations"`
	UnmapBytes                  int64 `json:"unmap_bytes"`
	InvalidWrOperations         int   `json:"invalid_wr_operations"`
	RdOperations                int   `json:"rd_operations"`
	WrOperations                int   `json:"wr_operations"`
}
