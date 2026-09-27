package permission

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type Permission uint64

const (
	CreateInstantInvite Permission = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel // Read messages
	SendMessages
	SendTTSMessages
	ManageMessages
	EmbedLinks
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVAD // Use voice activity
	ChangeNickname
	ManageNicknames
	ManageRoles // Manage permissions
	ManageWebhooks
	ManageEmojis
	UseApplicationCommands
	RequestToSpeak
	ManageEvents
	ManageThreads
	CreatePublicThreads
	CreatePrivateThreads
	UseExternalStickers
	SendMessagesInThreads
	UseEmbeddedActivities
	ModerateMembers
	ViewCreatorMonetizationAnalytics
	UseSoundboard
	CreateGuildExpressions
	CreateEvents
	UseExternalSounds
	SendVoiceMessages
	_
	_
	SendPolls
	UseExternalApps
	PinMessages
	BypassSlowmode
)

func HasPermissionRaw(permissions uint64, permission Permission) bool {
	return permissions&uint64(permission) == uint64(permission)
}

func BuildPermissions(permissions ...Permission) uint64 {
	var i uint64

	for _, permission := range permissions {
		i |= uint64(permission)
	}

	return i
}

func (p Permission) String() string {
	switch p {
	case CreateInstantInvite:
		return "Create Invite"
	case KickMembers:
		return "Kick, Approve and Reject Members"
	case BanMembers:
		return "Ban Members"
	case Administrator:
		return "Administrator"
	case ManageChannels:
		return "Manage Channels"
	case ManageGuild:
		return "Manage Server"
	case AddReactions:
		return "Add Reactions"
	case ViewAuditLog:
		return "View Audit Log"
	case PrioritySpeaker:
		return "Priority Speaker"
	case Stream:
		return "Video"
	case ViewChannel:
		return "View Channels"
	case SendMessages:
		return "Send Messages and Create Posts"
	case SendTTSMessages:
		return "Send Text-to-speech Messages"
	case ManageMessages:
		return "Manage Messages"
	case EmbedLinks:
		return "Embed Links"
	case AttachFiles:
		return "Attach Files"
	case ReadMessageHistory:
		return "Read Message History"
	case MentionEveryone:
		return "Mention @everyone, @here and All Roles"
	case UseExternalEmojis:
		return "Use External Emojis"
	case ViewGuildInsights:
		return "View Guild Insights"
	case Connect:
		return "Connect"
	case Speak:
		return "Speak"
	case MuteMembers:
		return "Mute Members"
	case DeafenMembers:
		return "Deafen Members"
	case MoveMembers:
		return "Move Members"
	case UseVAD:
		return "Use Voice Activity"
	case ChangeNickname:
		return "Change Nickname"
	case ManageNicknames:
		return "Manage Nicknames"
	case ManageRoles:
		return "Manage Roles"
	case ManageWebhooks:
		return "Manage Webhooks"
	case ManageEmojis:
		return "Manage Expressions"
	case UseApplicationCommands:
		return "Use Application Commands"
	case RequestToSpeak:
		return "Request to Speak"
	case ManageEvents:
		return "Manage Events"
	case ManageThreads:
		return "Manage Threads and Posts"
	case CreatePublicThreads:
		return "Create Public Threads"
	case CreatePrivateThreads:
		return "Create Private Threads"
	case UseExternalStickers:
		return "Use External Stickers"
	case SendMessagesInThreads:
		return "Send Messages in Threads and Posts"
	case UseEmbeddedActivities:
		return "Use Activities"
	case ModerateMembers:
		return "Time out members"
	case ViewCreatorMonetizationAnalytics:
		return "View Creator Monetization Analytics"
	case UseSoundboard:
		return "Use Soundboard"
	case CreateGuildExpressions:
		return "Create Expressions"
	case CreateEvents:
		return "Create Events"
	case UseExternalSounds:
		return "Use External Sounds"
	case SendVoiceMessages:
		return "Send Voice Messages"
	case SendPolls:
		return "Create Polls"
	case UseExternalApps:
		return "Use External Apps"
	case PinMessages:
		return "Pin Messages"
	case BypassSlowmode:
		return "Bypass Slowmode"
	default:
		return "Unknown Permission"
	}
}

func (p Permission) ChannelName(channelType channel.ChannelType) string {
	if channelType == channel.ChannelTypeGuildCategory {
		if p == ManageRoles {
			return "Manage Permissions"
		}

		return p.String()
	}

	if channelType.IsPostBased() {
		switch p {
		case SendMessages:
			return "Create Posts"
		case SendMessagesInThreads:
			return "Send Messages in Posts"
		case ManageThreads:
			return "Manage Posts"
		case ReadMessageHistory:
			return "Read Post History"
		}
	} else {
		switch p {
		case SendMessages:
			return "Send Messages"
		case SendMessagesInThreads:
			return "Send Messages in Threads"
		case ManageThreads:
			return "Manage Threads"
		}
	}

	switch p {
	case ViewChannel:
		return "View Channel"
	case ManageChannels:
		return "Manage Channel"
	case ManageRoles:
		return "Manage Permissions"
	default:
		return p.String()
	}
}

var AllPermissions = []Permission{
	CreateInstantInvite,
	KickMembers,
	BanMembers,
	Administrator,
	ManageChannels,
	ManageGuild,
	AddReactions,
	ViewAuditLog,
	PrioritySpeaker,
	Stream,
	ViewChannel,
	SendMessages,
	SendTTSMessages,
	ManageMessages,
	EmbedLinks,
	AttachFiles,
	ReadMessageHistory,
	MentionEveryone,
	UseExternalEmojis,
	ViewGuildInsights,
	Connect,
	Speak,
	MuteMembers,
	DeafenMembers,
	MoveMembers,
	UseVAD,
	ChangeNickname,
	ManageNicknames,
	ManageRoles,
	ManageWebhooks,
	ManageEmojis,
	UseApplicationCommands,
	RequestToSpeak,
	ManageEvents,
	ManageThreads,
	CreatePublicThreads,
	CreatePrivateThreads,
	UseExternalStickers,
	SendMessagesInThreads,
	UseEmbeddedActivities,
	ModerateMembers,
	ViewCreatorMonetizationAnalytics,
	UseSoundboard,
	CreateGuildExpressions,
	CreateEvents,
	UseExternalSounds,
	SendVoiceMessages,
	SendPolls,
	UseExternalApps,
	PinMessages,
	BypassSlowmode,
}
