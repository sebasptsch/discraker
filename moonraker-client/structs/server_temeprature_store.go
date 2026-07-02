package structs

type ServerTemperatureStore map[string]ServerTemperatureStoreSensor

type ServerTemperatureStoreSensor struct {
	Temperatures []int `json:"temperatures"`
	Targets      []int `json:"targets"`
	Speeds       []int `json:"speeds"`
}

type ServerTemperatureStoreParams struct {
	IncludeMonitors bool `json:"include_monitors"`
}
