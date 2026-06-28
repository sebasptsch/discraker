package structs

type ServerFilesPostDirectoryParams struct {
	Path string `json:"path"`
}

type ServerFilesPostDirectory struct {
	Item   ItemDetails `json:"item"`
	Action string      `json:"action"`
}

type ItemDetails struct {
	Path        string  `json:"path"`
	Root        string  `json:"root"`
	Modified    float64 `json:"modified"`
	Size        int     `json:"size"`
	Permissions string  `json:"permissions"`
}

type DestinationItem ItemDetails
