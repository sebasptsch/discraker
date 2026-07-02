package structs

type ServerFilesDeleteDirectoryParams struct {
	Path  string `json:"path"`
	Force bool   `json:"force"`
}

type ServerFilesDeleteDirectory struct {
	Item   ItemDetails `json:"item"`
	Action string      `json:"action"`
}
