package structs

type MachinePeripheralsVideo struct {
	V4L2Devices      []V4L2Device      `json:"v4l2_devices"`
	LibcameraDevices []LibcameraDevice `json:"libcamera_devices"`
}

type V4L2Device struct {
	DeviceName     string     `json:"device_name"`
	DevicePath     string     `json:"device_path"`
	CameraName     string     `json:"camera_name"`
	DriverName     string     `json:"driver_name"`
	AltName        string     `json:"alt_name"`
	HardwareBus    string     `json:"hardware_bus"`
	Capabilities   []string   `json:"capabilities"`
	Version        string     `json:"version"`
	PathByHardware string     `json:"path_by_hardware"`
	PathByID       string     `json:"path_by_id"`
	USBLocation    string     `json:"usb_location"`
	Modes          []V4L2Mode `json:"modes"`
}

type V4L2Mode struct {
	Format      string   `json:"format"`
	Description string   `json:"description"`
	Flags       []string `json:"flags"`
	Resolutions []string `json:"resolutions"`
}

type LibcameraDevice struct {
	LibcameraID string          `json:"libcamera_id"`
	Model       string          `json:"model"`
	Modes       []LibcameraMode `json:"modes"`
}

type LibcameraMode struct {
	Format      string `json:"format"`
	Resolutions string `json:"resolutions"`
}
