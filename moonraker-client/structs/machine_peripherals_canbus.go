package structs

type MachinePeripheralsCanbusParams struct {
	Interface string `json:"interface"`
}

type MachinePeripheralsCanbus struct {
	CANUUIDs []CANUUID `json:"can_uuids"`
}

type CANUUID struct {
	UUID        string `json:"uuid"`
	Application string `json:"application"`
}
