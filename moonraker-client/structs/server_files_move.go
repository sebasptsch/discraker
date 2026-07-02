package structs

type ServerFilesMoveParams struct {
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

type ServerFilesMove struct {
	Item       DestinationItem `json:"item"`
	SourceItem `json:"source_item"`
	Action     string `json:"action"`
}

type SourceItem struct {
	Root string `json:"root"`
	Path string `json:"path"`
}
