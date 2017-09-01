package messenger

type ContentType string

const (
	ContentTypeText     ContentType = "text"
	ContentTypeLocation ContentType = "location"
)

type SendMessage struct {
	Text         string       `json:"text,omitempty"`
	Attachment   *Attachment  `json:"attachment,omitempty"`
	QuickReplies []QuickReply `json:"quick_replies,omitempty"`
	Metadata     string       `json:"metadata,omitempty"`
}

type QuickReply struct {
	ContentType ContentType `json:"content_type"`
	Title       string      `json:"title,omitempty"`
	Payload     string      `json:"payload,omitempty"`
	ImageURL    string      `json:"image_url,omitempty"`
}

// Recipient describes the person who will receive the message
// Either ID or PhoneNumber has to be set
type Recipient struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// NotificationType describes the behavior phone will execute after receiving the message
type NotificationType string

const (
	// NotificationTypeRegular will emit a sound/vibration and a phone notification
	NotificationTypeRegular NotificationType = "REGULAR"
	// NotificationTypeSilentPush will just emit a phone notification
	NotificationTypeSilentPush NotificationType = "SILENT_PUSH"
	// NotificationTypeNoPush will not emit sound/vibration nor a phone notification
	NotificationTypeNoPush NotificationType = "NO_PUSH"
)

type MessageQuery struct {
	Recipient        Recipient        `json:"recipient"`
	Message          *SendMessage     `json:"message,omitempty"`
	NotificationType NotificationType `json:"notification_type,omitempty"`
	Action           SenderAction     `json:"sender_action,omitempty"`
}