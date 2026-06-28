package structs

type ServerFilesListParams struct {
	Root string `json:"root"`
}

type ServerFilesList []FileInfo

type FileInfo struct {
	Path        string  `json:"path"`
	Modified    float64 `json:"modified"`
	Size        int     `json:"size"`
	Permissions string  `json:"permissions"`
}
