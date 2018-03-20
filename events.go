package messenger

// UpstreamEvent represents messenger's incoming format
type UpstreamEvent struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

// Event represents a Webhook postback event.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type Event struct {
	ID   string `json:"id"`
	Time int64  `json:"time"`
}

// MessageOpts contains information common to all message events.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageOpts struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64 `json:"timestamp"`
}

// Entry represents an element of the incoming data
// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/#entry
type Entry struct {
	MessageOpts
	Message              *MessageEcho          `json:"message,omitempty"`
	Delivery             *Delivery             `json:"delivery,omitempty"`
	Postback             *Postback             `json:"postback,omitempty"`
	Optin                *Optin                `json:"optin,empty"`
	Read                 *Read                 `json:"read,omitempty"`
	AppRoles             *AppRoles             `json:"app_roles,omitempty"`
	PassThreadControl    *PassThreadControl    `json:"pass_thread_control,omitempty"`
	TakeThreadControl    *TakeThreadControl    `json:"take_thread_control,omitempty"`
	RequestThreadControl *RequestThreadControl `json:"request_thread_control,omitempty"`
	Standby              []Entry               `json:"standby,omitempty"`
}

// MessageEvent encapsulates common info plus the specific type of callback
// being received.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageEvent struct {
	Event
	Messaging []Entry `json:"messaging"`
}

// AppRole is a specific type for AppRoles
type AppRole string

// Valid appRoles
const (
	AppRolePrimaryReceiver AppRole = "primary_receiver"
	AppRolePersistentMenu  AppRole = "persistent_menu"
)

// RequestThreadControl This event will be sent when a page admin changes the role of your application
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/request-thread-control
type RequestThreadControl struct {
	RequestedOwnerAppID int64  `json:"requested_owner_app_id"`
	Metadata            string `json:"metadata"`
}

// PassThreadControl represents a thread ownership pass event
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/pass-thread-control
type PassThreadControl struct {
	NewOwnerAppID int64  `json:"new_owner_app_id"`
	Metadata      string `json:"metadata"`
}

// TakeThreadControl represents a thread ownership take event
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/take-thread-control
type TakeThreadControl struct {
	PreviousOwnerAppID int64  `json:"previous_owner_app_id"`
	Metadata           string `json:"metadata"`
}

// AppRoles represents a slice of AppRole for specific page(s)
// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_handovers#app_roles
type AppRoles map[string][]AppRole

// ReceivedMessage contains message specific information included with an echo callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type ReceivedMessage struct {
	ID          string             `json:"mid"`
	Text        string             `json:"text,omitempty"`
	Attachments []*Attachment      `json:"attachments,omitempty"`
	Seq         int                `json:"seq"`
	QuickReply  *QuickReplyPayload `json:"quick_reply,omitempty"`
	IsEcho      bool               `json:"is_echo,omitempty"`
	Metadata    *string            `json:"metadata,omitempty"`
}

// QuickReplyPayload contains content specific to a quick reply.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type QuickReplyPayload struct {
	Payload string
}

// Delivery contains information specific to a message delivered callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-delivered
type Delivery struct {
	MessageIDS []string `json:"mids"`
	Watermark  int64    `json:"watermark"`
	Seq        int      `json:"seq"`
}

// Postback contains content specific to a postback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type Postback struct {
	Payload string `json:"payload"`
}

// Optin contains information specific to Opt-In callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/optins
type Optin struct {
	Ref string `json:"ref"`
}

// Read contains data specific to message read callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-read
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

// MessageEcho contains information specific to an echo callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type MessageEcho struct {
	ReceivedMessage
	AppID int64 `json:"app_id,omitempty"`
}
