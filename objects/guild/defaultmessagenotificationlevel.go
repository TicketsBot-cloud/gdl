package guild

type DefaultMessageNotificationLevel int

const (
	DefaultMessageNotificationLevelAllMessages DefaultMessageNotificationLevel = iota
	DefaultMessageNotificationLevelOnlyMentions
)
