package messenger

import "encoding/json"

// AttachmentType is an attachment specific string
type AttachmentType string

// AttachmentType*  are defining attachment types
const (
	AttachmentTypeTemplate AttachmentType = "template"
	AttachmentTypeImage    AttachmentType = "image"
	AttachmentTypeVideo    AttachmentType = "video"
	AttachmentTypeAudio    AttachmentType = "audio"
	AttachmentTypeFile     AttachmentType = "file"
	AttachmentTypeLocation AttachmentType = "location"
)

// Attachment ...
type Attachment struct {
	Type    AttachmentType  `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// AttachmentPayload ...
type AttachmentPayload interface{}

// ImagePayload ...
type ImagePayload struct {
	Url string `json:"url"`
}

// VideoPayload ...
type VideoPayload struct {
	Url string `json:"url"`
}

// AudioPayload ...
type AudioPayload struct {
	Url string `json:"url"`
}

// FilePayload ...
type FilePayload struct {
	Url string `json:"url"`
}

// Coordinates ...
type Coordinates struct {
	Lat  float64
	Long float64
}

// Location ...
type Location struct {
	Coordinates Coordinates
}

// Resource ...
type Resource struct {
	URL      string `json:"url"`
	Reusable bool   `json:"is_reusable,omitempty"`
}

// ReusableAttachment ...
type ReusableAttachment struct {
	AttachmentID string `json:"attachment_id"`
}
