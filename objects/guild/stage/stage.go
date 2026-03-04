package stage

import "github.com/TicketsBot-cloud/gdl/objects"

type StagePrivacyLevel int

const (
	StagePrivacyLevelPublic    StagePrivacyLevel = 1 // Deprecated
	StagePrivacyLevelGuildOnly StagePrivacyLevel = 2
)

type StageInstance struct {
	Id                    uint64                    `json:"id,string"`
	GuildId               uint64                    `json:"guild_id,string"`
	ChannelId             uint64                    `json:"channel_id,string"`
	Topic                 string                    `json:"topic"`
	PrivacyLevel          StagePrivacyLevel         `json:"privacy_level"`
	DiscoverableDisabled  bool                      `json:"discoverable_disabled,omitempty"`
	GuildScheduledEventId objects.NullableSnowflake `json:"guild_scheduled_event_id"`
}
