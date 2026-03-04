package events

import "reflect"

type Event interface {
	Ready |
		Resumed |
		Reconnect |
		InvalidSession |
		ApplicationCommandPermissionsUpdate |
		AutoModerationRuleCreate |
		AutoModerationRuleUpdate |
		AutoModerationRuleDelete |
		AutoModerationActionExecution |
		ChannelCreate |
		ChannelUpdate |
		ChannelDelete |
		ChannelPinsUpdate |
		ThreadCreate |
		ThreadUpdate |
		ThreadDelete |
		ThreadListSync |
		ThreadMembersUpdate |
		ThreadMemberUpdate |
		EntitlementCreate |
		EntitlementUpdate |
		EntitlementDelete |
		GuildCreate |
		GuildUpdate |
		GuildDelete |
		GuildAuditLogEntryCreate |
		GuildBanAdd |
		GuildBanRemove |
		GuildEmojisUpdate |
		GuildStickersUpdate |
		GuildIntegrationsUpdate |
		GuildMemberAdd |
		GuildMemberRemove |
		GuildMemberUpdate |
		GuildMembersChunk |
		GuildRoleCreate |
		GuildRoleUpdate |
		GuildRoleDelete |
		GuildScheduledEventCreate |
		GuildScheduledEventUpdate |
		GuildScheduledEventDelete |
		GuildScheduledEventUserAdd |
		GuildScheduledEventUserRemove |
		GuildSoundboardSoundCreate |
		GuildSoundboardSoundUpdate |
		GuildSoundboardSoundDelete |
		GuildSoundboardSoundsUpdate |
		SoundboardSounds |
		IntegrationCreate |
		IntegrationUpdate |
		IntegrationDelete |
		InviteCreate |
		InviteDelete |
		MessageCreate |
		MessageUpdate |
		MessageDelete |
		MessageDeleteBulk |
		MessageReactionAdd |
		MessageReactionRemove |
		MessageReactionRemoveAll |
		MessageReactionRemoveEmoji |
		MessagePollVoteAdd |
		MessagePollVoteRemove |
		PresenceUpdate |
		StageInstanceCreate |
		StageInstanceUpdate |
		StageInstanceDelete |
		SubscriptionCreate |
		SubscriptionUpdate |
		SubscriptionDelete |
		TypingStart |
		UserUpdate |
		VoiceChannelEffectSend |
		VoiceServerUpdate |
		VoiceStateUpdate |
		WebhooksUpdate
}

var EventTypes = map[EventType]reflect.Type{
	READY:           reflect.TypeOf(Ready{}),
	RESUMED:         reflect.TypeOf(Resumed{}),
	RECONNECT:       reflect.TypeOf(Reconnect{}),
	INVALID_SESSION: reflect.TypeOf(InvalidSession{}),

	APPLICATION_COMMAND_PERMISSIONS_UPDATE: reflect.TypeOf(ApplicationCommandPermissionsUpdate{}),

	AUTO_MODERATION_RULE_CREATE:      reflect.TypeOf(AutoModerationRuleCreate{}),
	AUTO_MODERATION_RULE_UPDATE:      reflect.TypeOf(AutoModerationRuleUpdate{}),
	AUTO_MODERATION_RULE_DELETE:      reflect.TypeOf(AutoModerationRuleDelete{}),
	AUTO_MODERATION_ACTION_EXECUTION: reflect.TypeOf(AutoModerationActionExecution{}),

	CHANNEL_CREATE:      reflect.TypeOf(ChannelCreate{}),
	CHANNEL_UPDATE:      reflect.TypeOf(ChannelUpdate{}),
	CHANNEL_DELETE:      reflect.TypeOf(ChannelDelete{}),
	CHANNEL_PINS_UPDATE: reflect.TypeOf(ChannelPinsUpdate{}),

	THREAD_CREATE:         reflect.TypeOf(ThreadCreate{}),
	THREAD_UPDATE:         reflect.TypeOf(ThreadUpdate{}),
	THREAD_DELETE:         reflect.TypeOf(ThreadDelete{}),
	THREAD_LIST_SYNC:      reflect.TypeOf(ThreadListSync{}),
	THREAD_MEMBER_UPDATE:  reflect.TypeOf(ThreadMemberUpdate{}),
	THREAD_MEMBERS_UPDATE: reflect.TypeOf(ThreadMembersUpdate{}),

	ENTITLEMENT_CREATE: reflect.TypeOf(EntitlementCreate{}),
	ENTITLEMENT_UPDATE: reflect.TypeOf(EntitlementUpdate{}),
	ENTITLEMENT_DELETE: reflect.TypeOf(EntitlementDelete{}),

	GUILD_CREATE:                      reflect.TypeOf(GuildCreate{}),
	GUILD_UPDATE:                      reflect.TypeOf(GuildUpdate{}),
	GUILD_DELETE:                      reflect.TypeOf(GuildDelete{}),
	GUILD_AUDIT_LOG_ENTRY_CREATE:      reflect.TypeOf(GuildAuditLogEntryCreate{}),
	GUILD_BAN_ADD:                     reflect.TypeOf(GuildBanAdd{}),
	GUILD_BAN_REMOVE:                  reflect.TypeOf(GuildBanRemove{}),
	GUILD_EMOJIS_UPDATE:               reflect.TypeOf(GuildEmojisUpdate{}),
	GUILD_STICKERS_UPDATE:             reflect.TypeOf(GuildStickersUpdate{}),
	GUILD_INTEGRATIONS_UPDATE:         reflect.TypeOf(GuildIntegrationsUpdate{}),
	GUILD_MEMBER_ADD:                  reflect.TypeOf(GuildMemberAdd{}),
	GUILD_MEMBER_REMOVE:               reflect.TypeOf(GuildMemberRemove{}),
	GUILD_MEMBER_UPDATE:               reflect.TypeOf(GuildMemberUpdate{}),
	GUILD_MEMBERS_CHUNK:               reflect.TypeOf(GuildMembersChunk{}),
	GUILD_ROLE_CREATE:                 reflect.TypeOf(GuildRoleCreate{}),
	GUILD_ROLE_UPDATE:                 reflect.TypeOf(GuildRoleUpdate{}),
	GUILD_ROLE_DELETE:                 reflect.TypeOf(GuildRoleDelete{}),
	GUILD_SCHEDULED_EVENT_CREATE:      reflect.TypeOf(GuildScheduledEventCreate{}),
	GUILD_SCHEDULED_EVENT_UPDATE:      reflect.TypeOf(GuildScheduledEventUpdate{}),
	GUILD_SCHEDULED_EVENT_DELETE:      reflect.TypeOf(GuildScheduledEventDelete{}),
	GUILD_SCHEDULED_EVENT_USER_ADD:    reflect.TypeOf(GuildScheduledEventUserAdd{}),
	GUILD_SCHEDULED_EVENT_USER_REMOVE: reflect.TypeOf(GuildScheduledEventUserRemove{}),
	GUILD_SOUNDBOARD_SOUND_CREATE:     reflect.TypeOf(GuildSoundboardSoundCreate{}),
	GUILD_SOUNDBOARD_SOUND_UPDATE:     reflect.TypeOf(GuildSoundboardSoundUpdate{}),
	GUILD_SOUNDBOARD_SOUND_DELETE:     reflect.TypeOf(GuildSoundboardSoundDelete{}),
	GUILD_SOUNDBOARD_SOUNDS_UPDATE:    reflect.TypeOf(GuildSoundboardSoundsUpdate{}),
	SOUNDBOARD_SOUNDS:                 reflect.TypeOf(SoundboardSounds{}),

	INTEGRATION_CREATE: reflect.TypeOf(IntegrationCreate{}),
	INTEGRATION_UPDATE: reflect.TypeOf(IntegrationUpdate{}),
	INTEGRATION_DELETE: reflect.TypeOf(IntegrationDelete{}),

	INVITE_CREATE: reflect.TypeOf(InviteCreate{}),
	INVITE_DELETE: reflect.TypeOf(InviteDelete{}),

	MESSAGE_CREATE:                reflect.TypeOf(MessageCreate{}),
	MESSAGE_UPDATE:                reflect.TypeOf(MessageUpdate{}),
	MESSAGE_DELETE:                reflect.TypeOf(MessageDelete{}),
	MESSAGE_DELETE_BULK:           reflect.TypeOf(MessageDeleteBulk{}),
	MESSAGE_REACTION_ADD:          reflect.TypeOf(MessageReactionAdd{}),
	MESSAGE_REACTION_REMOVE:       reflect.TypeOf(MessageReactionRemove{}),
	MESSAGE_REACTION_REMOVE_ALL:   reflect.TypeOf(MessageReactionRemoveAll{}),
	MESSAGE_REACTION_REMOVE_EMOJI: reflect.TypeOf(MessageReactionRemoveEmoji{}),
	MESSAGE_POLL_VOTE_ADD:         reflect.TypeOf(MessagePollVoteAdd{}),
	MESSAGE_POLL_VOTE_REMOVE:      reflect.TypeOf(MessagePollVoteRemove{}),

	PRESENCE_UPDATE: reflect.TypeOf(PresenceUpdate{}),

	STAGE_INSTANCE_CREATE: reflect.TypeOf(StageInstanceCreate{}),
	STAGE_INSTANCE_UPDATE: reflect.TypeOf(StageInstanceUpdate{}),
	STAGE_INSTANCE_DELETE: reflect.TypeOf(StageInstanceDelete{}),

	SUBSCRIPTION_CREATE: reflect.TypeOf(SubscriptionCreate{}),
	SUBSCRIPTION_UPDATE: reflect.TypeOf(SubscriptionUpdate{}),
	SUBSCRIPTION_DELETE: reflect.TypeOf(SubscriptionDelete{}),

	TYPING_START: reflect.TypeOf(TypingStart{}),

	USER_UPDATE: reflect.TypeOf(UserUpdate{}),

	VOICE_CHANNEL_EFFECT_SEND: reflect.TypeOf(VoiceChannelEffectSend{}),
	VOICE_STATE_UPDATE:        reflect.TypeOf(VoiceStateUpdate{}),
	VOICE_SERVER_UPDATE:       reflect.TypeOf(VoiceServerUpdate{}),

	WEBHOOKS_UPDATE: reflect.TypeOf(WebhooksUpdate{}),
}
