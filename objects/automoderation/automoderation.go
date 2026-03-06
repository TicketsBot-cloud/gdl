package automoderation

import "github.com/TicketsBot-cloud/gdl/utils"

type EventType int

const (
	EventTypeMessageSend EventType = iota + 1
	EventTypeMemberUpdate
)

type TriggerType int

const (
	TriggerTypeKeyword TriggerType = iota + 1
	_
	TriggerTypeSpam
	TriggerTypeKeywordPreset
	TriggerTypeMentionSpam
	TriggerTypeMemberProfile
)

type KeywordPresetType int

const (
	KeywordPresetTypeProfanity KeywordPresetType = iota + 1
	KeywordPresetTypeSexualContent
	KeywordPresetTypeSlurs
)

type ActionType int

const (
	ActionTypeBlockMessage ActionType = iota + 1
	ActionTypeSendAlertMessage
	ActionTypeTimeout
	ActionTypeBlockMemberInteraction
)

type TriggerMetadata struct {
	KeywordFilter                []string            `json:"keyword_filter"`
	RegexPatterns                []string            `json:"regex_patterns"`
	Presets                      []KeywordPresetType `json:"presets"`
	AllowList                    []string            `json:"allow_list"`
	MentionTotalLimit            int                 `json:"mention_total_limit"`
	MentionRaidProtectionEnabled bool                `json:"mention_raid_protection_enabled"`
}

type ActionMetadata struct {
	ChannelId       *uint64 `json:"channel_id,string,omitempty"`
	DurationSeconds *int    `json:"duration_seconds,omitempty"`
	CustomMessage   *string `json:"custom_message,omitempty"`
}

type Action struct {
	Type     ActionType      `json:"type"`
	Metadata *ActionMetadata `json:"metadata,omitempty"`
}

type Rule struct {
	Id              uint64                  `json:"id,string"`
	GuildId         uint64                  `json:"guild_id,string"`
	Name            string                  `json:"name"`
	CreatorId       uint64                  `json:"creator_id,string"`
	EventType       EventType               `json:"event_type"`
	TriggerType     TriggerType             `json:"trigger_type"`
	TriggerMetadata TriggerMetadata         `json:"trigger_metadata"`
	Actions         []Action                `json:"actions"`
	Enabled         bool                    `json:"enabled"`
	ExemptRoles     utils.Uint64StringSlice `json:"exempt_roles"`
	ExemptChannels  utils.Uint64StringSlice `json:"exempt_channels"`
}
