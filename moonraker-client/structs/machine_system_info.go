package structs

type MachineSystemInfo struct {
	SystemInfo SystemInfo `json:"system_info"`
}

type SystemInfo struct {
	Python            PythonInfo         `json:"python"`
	CPUInfo           CPUInfo            `json:"cpu_info"`
	SDInfo            SDInfo             `json:"sd_info"`
	Distribution      DistributionInfo   `json:"distribution"`
	Virtualization    VirtualizationInfo `json:"virtualization"`
	Network           NetworkInfo        `json:"network"`
	Canbus            CanbusInfo         `json:"canbus"`
	Provider          string             `json:"provider"`
	AvailableServices []string           `json:"available_services"`
	ServiceState      ServiceStateInfo   `json:"service_state"`
	InstanceIdsInfo   InstanceIdsInfo    `json:"instance_ids"`
}

type PythonInfo struct {
	Version       []string `json:"version"`
	VersionString string   `json:"version_string"`
}

type CPUInfo struct {
	CPUCount     int    `json:"cpu_count"`
	Bits         string `json:"bits"`
	CPUDesc      string `json:"cpu_desc"`
	SerialNumber string `json:"serial_number"`
	HardwareDesc string `json:"hardware_desc"`
	Model        string `json:"model"`
	TotalMemory  int    `json:"total_memory"`
	MemoryUnits  string `json:"memory_units"`
}

type SDInfo struct {
	ManufacturerID   string `json:"manufacturer_id"`
	Manufacturer     string `json:"manufacturer"`
	OEMID            string `json:"oem_id"`
	ProductName      string `json:"product_name"`
	ProductRevision  string `json:"product_revision"`
	SerialNumber     string `json:"serial_number"`
	ManufacturerDate string `json:"manufacturer_date"`
	Capacity         string `json:"capacity"`
	TotalBytes       int    `json:"total_bytes"`
}

type DistributionInfo struct {
	Name         string       `json:"name"`
	ID           string       `json:"id"`
	Like         string       `json:"like"`
	Codename     string       `json:"codename"`
	Version      string       `json:"version"`
	VersionParts VersionParts `json:"version_parts"`
	ReleaseInfo  map[string]string
}

type VersionParts struct {
	Major   string `json:"major"`
	Minor   string `json:"minor"`
	Release string `json:"release"`
}

type VirtualizationInfo struct {
	VirtType string `json:"virt_type"`
	VirtID   string `json:"virt_id"`
}

type NetworkInfo map[string]NetworkInterface

type NetworkInterface struct {
	IPAddresses []IPAddress `json:"ip_addresses"`
	MACAddress  string      `json:"mac_address"`
}

type IPAddress struct {
	Address     string `json:"address"`
	Family      string `json:"family"`
	IsLinkLocal bool   `json:"is_link_local"`
}

type CanbusInfo map[string]CanbusInterface

type CanbusInterface struct {
	TxQueueLen int    `json:"tx_queue_len"`
	Bitrate    int    `json:"bitrate"`
	Driver     string `json:"driver"`
}

type ServiceStateInfo map[string]UnitStatus

type UnitStatus struct {
	ActiveState string `json:"active_state"`
	SubState    string `json:"sub_state"`
}

type InstanceIdsInfo struct {
	Klipper   string `json:"klipper"`
	Moonraker string `json:"moonraker"`
}
