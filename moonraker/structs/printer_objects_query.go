package structs

type PrinterObjectsQueryParams struct {
	Objects any `json:"objects"`
}

type PrinterObjectsQuery struct {
	EventTime float32 `json:"eventtime"`
	Status    any     `json:"status"`
}
