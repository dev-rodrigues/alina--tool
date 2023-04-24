package domain

type Service struct {
	PodID     string `json:"pod_id"`
	TotalPods int    `json:"total_pods"`
}

type ServiceDetails struct {
	PodID         string  `json:"pod_id"`
	TotalPods     int     `json:"total_pods"`
	Name          string  `json:"name"`
	State         string  `json:"state"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent uint64  `json:"memory_percent"`
	CreatedAt     int64   `json:"created_at"`
}
