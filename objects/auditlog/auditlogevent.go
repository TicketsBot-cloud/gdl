package auditlog

type AuditLogEvent int

const (
	AuditLogEventGuildUpdate AuditLogEvent = 1
)

const (
	AuditLogEventChannelCreate AuditLogEvent = iota + 10
	AuditLogEventChannelUpdate
	AuditLogEventChannelDelete
	AuditLogEventChannelOverwriteCreate
	AuditLogEventChannelOverwriteUpdate
	AuditLogEventChannelOverwriteDelete
)

const (
	AuditLogEventMemberKick AuditLogEvent = iota + 20
	AuditLogEventMemberPrune
	AuditLogEventMemberBanAdd
	AuditLogEventMemberBanRemove
	AuditLogEventMemberUpdate
	AuditLogEventMemberRoleUpdate
	AuditLogEventMemberMove
	AuditLogEventMemberDisconnect
	AuditLogEventBotAdd
)

const (
	AuditLogEventRoleCreate AuditLogEvent = iota + 30
	AuditLogEventRoleUpdate
	AuditLogEventRoleDelete
)

const (
	AuditLogEventInviteCreate AuditLogEvent = iota + 40
	AuditLogEventInviteUpdate
	AuditLogEventInviteDelete
)

const (
	AuditLogEventWebhookCreate AuditLogEvent = iota + 50
	AuditLogEventWebhookUpdate
	AuditLogEventWebhookDelete
)

const (
	AuditLogEventEmojiCreate AuditLogEvent = iota + 60
	AuditLogEventEmojiUpdate
	AuditLogEventEmojiDelete
)

const (
	AuditLogEventMessageDelete AuditLogEvent = iota + 72
	AuditLogEventMessageBulkDelete
	AuditLogEventMessagePin
	AuditLogEventMessageUnpin
)

const (
	AuditLogEventIntegrationCreate AuditLogEvent = iota + 80
	AuditLogEventIntegrationUpdate
	AuditLogEventIntegrationDelete
	AuditLogEventStageInstanceCreate
	AuditLogEventStageInstanceUpdate
	AuditLogEventStageInstanceDelete
)

const (
	AuditLogEventStickerCreate AuditLogEvent = iota + 90
	AuditLogEventStickerUpdate
	AuditLogEventStickerDelete
)

const (
	AuditLogEventGuildScheduledEventCreate AuditLogEvent = iota + 100
	AuditLogEventGuildScheduledEventUpdate
	AuditLogEventGuildScheduledEventDelete
)

const (
	AuditLogEventThreadCreate AuditLogEvent = iota + 110
	AuditLogEventThreadUpdate
	AuditLogEventThreadDelete
)

const (
	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121
)

const (
	AuditLogEventSoundboardSoundCreate AuditLogEvent = iota + 130
	AuditLogEventSoundboardSoundUpdate
	AuditLogEventSoundboardSoundDelete
)

const (
	AuditLogEventAutoModerationRuleCreate AuditLogEvent = iota + 140
	AuditLogEventAutoModerationRuleUpdate
	AuditLogEventAutoModerationRuleDelete
	AuditLogEventAutoModerationBlockMessage
	AuditLogEventAutoModerationFlagToChannel
	AuditLogEventAutoModerationUserCommunicationDis
)

const (
	AuditLogEventCreatorMonetizationRequestCreated AuditLogEvent = iota + 150
	AuditLogEventCreatorMonetizationTermsAccepted
)

const (
	AuditLogEventOnboardingPromptCreate AuditLogEvent = iota + 163
	AuditLogEventOnboardingPromptUpdate
	AuditLogEventOnboardingPromptDelete
	AuditLogEventOnboardingCreate
	AuditLogEventOnboardingUpdate
)

const (
	AuditLogEventHomeSettingsCreate AuditLogEvent = iota + 190
	AuditLogEventHomeSettingsUpdate
)
