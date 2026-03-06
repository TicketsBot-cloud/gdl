package message

type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = iota + 1
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	_
	MessageActivityTypeJoinRequest
)
