package structs

type ServerConfig struct {
	Config any          `json:"config"`
	Orig   any          `json:"orig"`
	Files  []ServerFile `json:"files"`
}

type ServerFile struct {
	Filename string   `json:"filename"`
	Sections []string `json:"sections"`
}
