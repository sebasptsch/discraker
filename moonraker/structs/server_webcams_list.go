package structs

type WebcamListEntry struct {
	Name           string `json:"name"`
	Location       string `json:"location"`
	Service        string `json:"service"`
	Enabled        bool   `json:"enabled"`
	Icon           string `json:"icon"`
	TargetFPS      int    `json:"target_fps"`
	TargetFPSIdle  int    `json:"target_fps_idle"`
	StreamURL      string `json:"stream_url"`
	SnapshotURL    string `json:"snapshot_url"`
	FlipHorizontal bool   `json:"flip_horizontal"`
	FlipVertical   bool   `json:"flip_vertical"`
	Rotation       int    `json:"rotation"`
	AspectRatio    string `json:"aspect_ratio"`
	// ExtraData
	Source string `json:"source"`
	UID    string `json:"uid"`
}

type WebcamList struct {
	Webcams []WebcamListEntry `json:"webcams"`
}
