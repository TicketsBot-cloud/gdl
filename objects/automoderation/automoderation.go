package automoderation

type EventType int

const (
	EventTypeMessageSend  EventType = 1
	EventTypeMemberUpdate EventType = 2
)

type TriggerType int

const (
	TriggerTypeKeyword       TriggerType = 1
	TriggerTypeSpam          TriggerType = 3
	TriggerTypeKeywordPreset TriggerType = 4
	TriggerTypeMentionSpam   TriggerType = 5
	TriggerTypeMemberProfile TriggerType = 6
)

type ActionType int

const (
	ActionTypeBlockMessage            ActionType = 1
	ActionTypeSendAlertMessage        ActionType = 2
	ActionTypeTimeout                 ActionType = 3
	ActionTypeBlockMemberInteraction  ActionType = 4
)

type TriggerMetadata struct {
	KeywordFilter                []string `json:"keyword_filter"`
	RegexPatterns                []string `json:"regex_patterns"`
	Presets                      []int    `json:"presets"`
	AllowList                    []string `json:"allow_list"`
	MentionTotalLimit            int      `json:"mention_total_limit"`
	MentionRaidProtectionEnabled bool     `json:"mention_raid_protection_enabled"`
}

type ActionMetadata struct {
	ChannelId       uint64  `json:"channel_id,string"`
	DurationSeconds int     `json:"duration_seconds"`
	CustomMessage   *string `json:"custom_message,omitempty"`
}

type Action struct {
	Type     ActionType      `json:"type"`
	Metadata *ActionMetadata `json:"metadata,omitempty"`
}

type Rule struct {
	Id              uint64          `json:"id,string"`
	GuildId         uint64          `json:"guild_id,string"`
	Name            string          `json:"name"`
	CreatorId       uint64          `json:"creator_id,string"`
	EventType       EventType       `json:"event_type"`
	TriggerType     TriggerType     `json:"trigger_type"`
	TriggerMetadata TriggerMetadata `json:"trigger_metadata"`
	Actions         []Action        `json:"actions"`
	Enabled         bool            `json:"enabled"`
	ExemptRoles     []uint64        `json:"exempt_roles"`
	ExemptChannels  []uint64        `json:"exempt_channels"`
}
