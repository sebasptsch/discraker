package structs

type MachineSudoInfoParams struct {
	CheckAccess bool `json:"check_access"`
}

type MachineSudoInfo struct {
	SudoAccess      bool     `json:"sudo_access"`
	LinuxUser       string   `json:"linux_user"`
	SudoRequested   bool     `json:"sudo_requested"`
	RequestMessages []string `json:"request_messages"`
}
