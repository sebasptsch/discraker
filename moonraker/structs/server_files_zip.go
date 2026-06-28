package structs

type ServerFilesZipParams struct {
	Dest      string   `json:"dest"`
	Items     []string `json:"items"`
	StoreOnly bool     `json:"store_only"`
}

type ServerFilesZip struct {
	Destination DestinationItem `json:"destination"`
	Action      string          `json:"action"`
}
