package scheduledevent

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type PrivacyLevel int

const (
	PrivacyLevelGuildOnly PrivacyLevel = 2
)

type EntityType int

const (
	EntityTypeStageInstance EntityType = iota + 1
	EntityTypeVoice
	EntityTypeExternal
)

type Status int

const (
	StatusScheduled Status = iota + 1
	StatusActive
	StatusCompleted
	StatusCanceled
)

type EntityMetadata struct {
	Location *string `json:"location,omitempty"`
}

type GuildScheduledEvent struct {
	Id                 uint64          `json:"id,string"`
	GuildId            uint64          `json:"guild_id,string"`
	ChannelId          *uint64         `json:"channel_id,string"`
	CreatorId          *uint64         `json:"creator_id,string"`
	Name               string          `json:"name"`
	Description        *string         `json:"description,omitempty"`
	ScheduledStartTime time.Time       `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time      `json:"scheduled_end_time,omitempty"`
	PrivacyLevel       PrivacyLevel    `json:"privacy_level"`
	Status             Status          `json:"status"`
	EntityType         EntityType      `json:"entity_type"`
	EntityId           *uint64         `json:"entity_id,string"`
	EntityMetadata     *EntityMetadata `json:"entity_metadata,omitempty"`
	Creator            *user.User      `json:"creator,omitempty"`
	UserCount          *int            `json:"user_count,omitempty"`
	Image              *string         `json:"image,omitempty"`
}
