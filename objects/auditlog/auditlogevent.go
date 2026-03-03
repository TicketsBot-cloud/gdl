package auditlog

type AuditLogEvent int

const (
	AuditLogEventGuildUpdate            AuditLogEvent = 1
	AuditLogEventChannelCreate          AuditLogEvent = 10
	AuditLogEventChannelUpdate          AuditLogEvent = 11
	AuditLogEventChannelDelete          AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete AuditLogEvent = 15
	AuditLogEventMemberKick             AuditLogEvent = 20
	AuditLogEventMemberPrune            AuditLogEvent = 21
	AuditLogEventMemberBanAdd           AuditLogEvent = 22
	AuditLogEventMemberBanRemove        AuditLogEvent = 23
	AuditLogEventMemberUpdate           AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate       AuditLogEvent = 25
	AuditLogEventMemberMove             AuditLogEvent = 26
	AuditLogEventMemberDisconnect       AuditLogEvent = 27
	AuditLogEventBotAdd                 AuditLogEvent = 28
	AuditLogEventRoleCreate             AuditLogEvent = 30
	AuditLogEventRoleUpdate             AuditLogEvent = 31
	AuditLogEventRoleDelete             AuditLogEvent = 32
	AuditLogEventInviteCreate           AuditLogEvent = 40
	AuditLogEventInviteUpdate           AuditLogEvent = 41
	AuditLogEventInviteDelete           AuditLogEvent = 42
	AuditLogEventWebhookCreate          AuditLogEvent = 50
	AuditLogEventWebhookUpdate          AuditLogEvent = 51
	AuditLogEventWebhookDelete          AuditLogEvent = 52
	AuditLogEventEmojiCreate            AuditLogEvent = 60
	AuditLogEventEmojiUpdate            AuditLogEvent = 61
	AuditLogEventEmojiDelete            AuditLogEvent = 62
	AuditLogEventMessageDelete          AuditLogEvent = 72
	AuditLogEventMessageBulkDelete      AuditLogEvent = 73
	AuditLogEventMessagePin             AuditLogEvent = 74
	AuditLogEventMessageUnpin           AuditLogEvent = 75
	AuditLogEventIntegrationCreate      AuditLogEvent = 80
	AuditLogEventIntegrationUpdate      AuditLogEvent = 81
	AuditLogEventIntegrationDelete      AuditLogEvent = 82
	AuditLogEventStageInstanceCreate    AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate    AuditLogEvent = 84
	AuditLogEventStageInstanceDelete    AuditLogEvent = 85
	AuditLogEventStickerCreate          AuditLogEvent = 90
	AuditLogEventStickerUpdate          AuditLogEvent = 91
	AuditLogEventStickerDelete          AuditLogEvent = 92

	AuditLogEventGuildScheduledEventCreate AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete AuditLogEvent = 102

	AuditLogEventThreadCreate AuditLogEvent = 110
	AuditLogEventThreadUpdate AuditLogEvent = 111
	AuditLogEventThreadDelete AuditLogEvent = 112

	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121

	AuditLogEventSoundboardSoundCreate AuditLogEvent = 130
	AuditLogEventSoundboardSoundUpdate AuditLogEvent = 131
	AuditLogEventSoundboardSoundDelete AuditLogEvent = 132

	AuditLogEventAutoModerationRuleCreate           AuditLogEvent = 140
	AuditLogEventAutoModerationRuleUpdate           AuditLogEvent = 141
	AuditLogEventAutoModerationRuleDelete           AuditLogEvent = 142
	AuditLogEventAutoModerationBlockMessage         AuditLogEvent = 143
	AuditLogEventAutoModerationFlagToChannel        AuditLogEvent = 144
	AuditLogEventAutoModerationUserCommunicationDis AuditLogEvent = 145

	AuditLogEventCreatorMonetizationRequestCreated AuditLogEvent = 150
	AuditLogEventCreatorMonetizationTermsAccepted  AuditLogEvent = 151

	AuditLogEventOnboardingPromptCreate AuditLogEvent = 163
	AuditLogEventOnboardingPromptUpdate AuditLogEvent = 164
	AuditLogEventOnboardingPromptDelete AuditLogEvent = 165
	AuditLogEventOnboardingCreate       AuditLogEvent = 166
	AuditLogEventOnboardingUpdate       AuditLogEvent = 167

	AuditLogEventHomeSettingsCreate AuditLogEvent = 190
	AuditLogEventHomeSettingsUpdate AuditLogEvent = 191
)
