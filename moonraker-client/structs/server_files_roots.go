package structs

type ServerFilesRoots []RootInfo

type RootInfo struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Permissions string `json:"permissions"`
}
