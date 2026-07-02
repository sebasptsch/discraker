package structs

import "io"

type ServerFilesUploadParams struct {
	Root     string  `json:"root"`
	Path     *string `json:"path,omitempty"`
	Checksum *string `json:"checksum,omitempty"`
	Print    *bool   `json:"print,omitempty"` // false
	File     io.Reader
}

type ServerFilesUpload struct {
	Item         DestinationItem `json:"item"`
	PrintStarted bool            `json:"print_started"`
	PrintQueued  bool            `json:"print_queued"`
	Action       string          `json:"action"`
}
