package structs

type ServerGcodeStore struct {
	GcodeStore []ServerGcodeTrackingObject `json:"gcode_store"`
}

type ServerGcodeTrackingObject struct {
	Message string  `json:"message"`
	Time    float32 `json:"time"`
	Type    string  `json:"type"`
}

type ServerGcodeStoreParams struct {
	Count int `json:"count"`
}
