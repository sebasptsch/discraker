package structs

type ServerFilesDeleteParams struct {
	Path string `json:"path"`
}

type ServerFilesDelete struct {
	Item   ItemDetails `json:"item"`
	Action string      `json:"action"`
}
