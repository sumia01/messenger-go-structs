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

type Attachment struct {
	Type    AttachmentType  `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type AttachmentPayload interface{}

type ImagePayload struct {
	Url string `json:"url"`
}

type VideoPayload struct {
	Url string `json:"url"`
}

type AudioPayload struct {
	Url string `json:"url"`
}

type FilePayload struct {
	Url string `json:"url"`
}

type Coordinates struct {
	Lat  float64
	Long float64
}

type Location struct {
	Coordinates Coordinates
}

type Resource struct {
	URL      string `json:"url"`
	Reusable bool   `json:"is_reusable,omitempty"`
}

type ReusableAttachment struct {
	AttachmentID string `json:"attachment_id"`
}
