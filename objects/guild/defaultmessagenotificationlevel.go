package guild

type DefaultMessageNotificationLevel int

const (
	DefaultMessageNotificationLevelAllMessages  DefaultMessageNotificationLevel = 0
	DefaultMessageNotificationLevelOnlyMentions DefaultMessageNotificationLevel = 1
)
