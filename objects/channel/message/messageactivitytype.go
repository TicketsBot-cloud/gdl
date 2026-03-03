package message

type MessageActivityType int

const (
	MessageActivityTypeJoin        MessageActivityType = 1
	MessageActivityTypeSpectate    MessageActivityType = 2
	MessageActivityTypeListen      MessageActivityType = 3
	MessageActivityTypeJoinRequest MessageActivityType = 5
)
