package structs

type MachinePeripheralsUSB struct {
	USBDevices []USBDevice `json:"usb_devices"`
}

type USBDevice struct {
	BusNum       int    `json:"bus_num"`
	DeviceNum    int    `json:"device_num"`
	USBLocation  string `json:"usb_location"`
	VendorID     string `json:"vendor_id"`
	ProductID    string `json:"product_id"`
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
	Class        string `json:"class"`
	Subclass     string `json:"subclass"`
	Protocol     string `json:"protocol"`
	Description  string `json:"description"`
}
