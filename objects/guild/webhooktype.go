package guild

type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
	WebhookTypeApplication
)
