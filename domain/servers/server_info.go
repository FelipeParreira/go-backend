package servers

type SummaryInfo struct {
	Hostname    HostName        `json:"hostname"`
	CPULoad     BaseSummaryInfo `json:"cpu_load"`
	MemoryUsage BaseSummaryInfo `json:"memory_usage"`
	DiskUsage   BaseSummaryInfo `json:"disk_usage"`
}

type BaseSummaryInfo struct {
	Average float64 `json:"average"`
	Mode    float64 `json:"mode"`
	Median  float64 `json:"median"`
	Unit    string  `json:"unit"`
}

type Info struct {
	Hostname    HostName `json:"hostname"`
	CPULoad     BaseInfo `json:"cpu_load"`
	MemoryUsage BaseInfo `json:"memory_usage"`
	DiskUsage   BaseInfo `json:"disk_usage"`
}

type BaseInfo struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
