package messenger

// SenderAction is an action specific string
type SenderAction string

// SenderAction* are defining sender actions
const (
	SenderActionMarkSeen SenderAction = "mark_seen"
	//SenderActionTypingOn indicator is automatically turned off after 20 seconds
	SenderActionTypingOn  SenderAction = "typing_on"
	SenderActionTypingOff SenderAction = "typing_off"
)
