package message

type MessageType int

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
)

const (
	MessageTypeGuildDiscoveryDisqualified MessageType = iota + 14
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeApplicationCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
	MessageTypeContextMenuCommand
	MessageTypeAutoModerationAction
	MessageTypeRoleSubscriptionPurchase
	MessageTypeInteractionPremiumUpsell
	MessageTypeStageStart
	MessageTypeStageEnd
	MessageTypeStageSpeaker
)

const (
	MessageTypeStageTopic MessageType = iota + 31
	MessageTypeGuildApplicationPremiumSubscription
)

const (
	MessageTypeGuildIncidentAlertModeEnabled MessageType = iota + 36
	MessageTypeGuildIncidentAlertModeDisabled
	MessageTypeGuildIncidentReportRaid
	MessageTypeGuildIncidentReportFalseAlarm
)

const (
	MessageTypePurchaseNotification MessageType = iota + 44
)

const (
	MessageTypePollResult MessageType = iota + 46
)
