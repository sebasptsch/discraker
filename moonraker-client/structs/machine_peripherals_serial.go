package structs

type MachinePeripheralsSerial struct {
	SerialDevices []SerialDevice `json:"serial_devices"`
}

type SerialDevice struct {
	DeviceType     string `json:"device_type"`
	DevicePath     string `json:"device_path"`
	DeviceName     string `json:"device_name"`
	PathByHardware string `json:"path_by_hardware"`
	PathByID       string `json:"path_by_id"`
	USBLocation    string `json:"usb_location"`
}
