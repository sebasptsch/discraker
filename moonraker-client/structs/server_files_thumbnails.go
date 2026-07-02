package structs

type ThumbnailDetails struct {
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Size          int    `json:"size"`
	ThumbnailPath string `json:"thumbnail_path"`
}

type ServerFilesThumbnails []ThumbnailDetails
