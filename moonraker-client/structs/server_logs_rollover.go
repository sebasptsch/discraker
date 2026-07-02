package structs

type ServerRolloverLogs struct {
	RolledOver []string `json:"rolled_over"`
	Failed     []string `json:"failed"`
}

type ServerRolloverLogsParams struct {
	Application string `json:"application"`
}
