package structs

type AccessInfo struct {
	DefaultSource    string   `json:"default_source"`
	AvailableSources []string `json:"available_sources"`
	LoginRequired    bool     `json:"login_required"`
	Trusted          bool     `json:"trusted"`
}
