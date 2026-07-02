package structs

type ServerFilesCopyParams struct {
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

type ServerFilesCopy struct {
	Item   DestinationItem `json:"item"`
	Action string          `json:"action"`
}
