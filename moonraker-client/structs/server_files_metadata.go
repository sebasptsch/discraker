package structs

type ServerFilesMetadataParams struct {
	Filename string `json:"filename"`
}

type ServerFilesMetadata struct {
	Size                int             `json:"size"`
	Modified            float64         `json:"modified"`
	UUID                string          `json:"uuid"`
	FilePreprocessors   []string        `json:"file_preprocessors"`
	Slicer              string          `json:"slicer"`
	SlicerVersion       string          `json:"slicer_version"`
	GcodeStartByte      int             `json:"gcode_start_byte"`
	GcodeIntByte        int             `json:"gcode_int_byte"`
	ObjectHeight        float32         `json:"object_height"`
	EstimatedTime       float32         `json:"estimated_time"`
	NozzleDiameter      float32         `json:"nozzle_diameter"`
	LayerHeight         float32         `json:"layer_height"`
	FirstLayerHeight    float32         `json:"first_layer_height"`
	FirstLayerExtrTemp  float32         `json:"first_layer_extr_temp"`
	FirstLayerBedTemp   float32         `json:"first_layer_bed_temp"`
	ChamberTemp         float32         `json:"chamber_temp"`
	FilamentName        string          `json:"filament_name"`
	FilamentColors      []string        `json:"filament_colors"`
	ExtruderColors      []string        `json:"extruder_colors"`
	FilamentTemps       []int           `json:"filament_temps"`
	FilamentType        string          `json:"filament_type"`
	FilamentTotal       float32         `json:"filament_total"`
	FilamentChangeCount int             `json:"filament_change_count"`
	FilamentWeightTotal float32         `json:"filament_weight_total"`
	FilamentWeights     []float32       `json:"filament_weights"`
	PrinterVendor       string          `json:"printer_vendor"`
	PrinterModel        string          `json:"printer_model"`
	PrinterVariant      string          `json:"printer_variant"`
	ProfileVersion      string          `json:"profile_version"`
	MMUPrint            int             `json:"mmu_print"`
	ReferencedTools     []int           `json:"referenced_tools"`
	Thumbnails          []ThumbnailInfo `json:"thumbnails"`
	JobID               string          `json:"job_id"`
	PrintStartTime      float64         `json:"print_start_time"`
	Filename            string          `json:"filename"`
}

type ThumbnailInfo struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Size         int    `json:"size"`
	RelativePath string `json:"relative_path"`
}
