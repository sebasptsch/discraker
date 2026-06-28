package structs

type MachineProcStats struct {
	MoonrakerStats       []MoonrakerStats `json:"moonraker_stats"`
	ThrottledState       ThrottledState   `json:"throttled_state"`
	CPUTemp              float32          `json:"cpu_temp"`
	Network              NetworkUsage     `json:"network"`
	SystemCPUUsage       CPUUsage         `json:"system_cpu_usage"`
	SystemMemory         MemoryUsage      `json:"system_memory"`
	SystemUptime         float32          `json:"system_uptime"`
	WebsocketConnections int              `json:"websocket_connections"`
}

type MoonrakerStats struct {
	Time     float64 `json:"time"`
	CPUUsage float32 `json:"cpu_usage"`
	Memory   int     `json:"memory"`
	MemUnits string  `json:"mem_units"`
}

type ThrottledState struct {
	Bits  int      `json:"bits"`
	Flags []string `json:"flags"`
}

type NetworkUsage map[string]InterfaceUsage

type InterfaceUsage struct {
	Bandwidth float32 `json:"bandwidth"`
	RXBytes   int     `json:"rx_bytes"`
	TXBytes   int     `json:"tx_bytes"`
	RXPackets int     `json:"rx_packets"`
	TXPackets int     `json:"tx_packets"`
	RXErrs    int     `json:"rx_errs"`
	TXErrs    int     `json:"tx_errs"`
	RXDrop    int     `json:"rx_drop"`
	TXDrop    int     `json:"tx_drop"`
}

type CPUUsage map[string]float32

type MemoryUsage struct {
	Total     int `json:"total"`
	Available int `json:"available"`
	Used      int `json:"used"`
}
