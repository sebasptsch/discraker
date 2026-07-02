package structs

type ServerFilesGetDirectoryParams struct {
	Path string `json:"path"`
}

type ServerFilesGetDirectory struct {
	Dirs      []DirectoryInfo    `json:"dirs"`
	Files     []FilesGetFileInfo `json:"files"`
	DiskUsage `json:"disk_usage"`
	RootInfo  `json:"root_info"`
}

type DirectoryInfo struct {
	Modified    float64 `json:"modified"`
	Size        int     `json:"size"`
	Permissions string  `json:"permissions"`
	Dirname     string  `json:"dirname"`
}

type FilesGetFileInfo struct {
	Modified    float64 `json:"modified"`
	Size        int     `json:"size"`
	Permissions string  `json:"permissions"`
	Filename    string  `json:"filename"`
}

type DiskUsage struct {
	Free  int `json:"free"`
	Used  int `json:"used"`
	Total int `json:"total"`
}
