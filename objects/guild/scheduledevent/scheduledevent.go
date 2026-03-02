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

type RecurrenceFrequency int

const (
	RecurrenceFrequencyYearly  RecurrenceFrequency = iota
	RecurrenceFrequencyMonthly
	RecurrenceFrequencyWeekly
	RecurrenceFrequencyDaily
)

type RecurrenceWeekday int

const (
	RecurrenceWeekdayMonday    RecurrenceWeekday = iota
	RecurrenceWeekdayTuesday
	RecurrenceWeekdayWednesday
	RecurrenceWeekdayThursday
	RecurrenceWeekdayFriday
	RecurrenceWeekdaySaturday
	RecurrenceWeekdaySunday
)

type RecurrenceMonth int

const (
	RecurrenceMonthJanuary  RecurrenceMonth = iota + 1
	RecurrenceMonthFebruary
	RecurrenceMonthMarch
	RecurrenceMonthApril
	RecurrenceMonthMay
	RecurrenceMonthJune
	RecurrenceMonthJuly
	RecurrenceMonthAugust
	RecurrenceMonthSeptember
	RecurrenceMonthOctober
	RecurrenceMonthNovember
	RecurrenceMonthDecember
)

type RecurrenceNWeekday struct {
	N   int               `json:"n"`
	Day RecurrenceWeekday `json:"day"`
}

type RecurrenceRule struct {
	Start      time.Time            `json:"start"`
	End        *time.Time           `json:"end,omitempty"`
	Frequency  RecurrenceFrequency  `json:"frequency"`
	Interval   int                  `json:"interval"`
	ByWeekday  []RecurrenceWeekday  `json:"by_weekday,omitempty"`
	ByNWeekday []RecurrenceNWeekday `json:"by_n_weekday,omitempty"`
	ByMonth    []RecurrenceMonth    `json:"by_month,omitempty"`
	ByMonthDay []int                `json:"by_month_day,omitempty"`
	ByYearDay  []int                `json:"by_year_day,omitempty"`
	Count      *int                 `json:"count,omitempty"`
}

type EntityMetadata struct {
	Location *string `json:"location,omitempty"`
}

type GuildScheduledEvent struct {
	Id                 uint64          `json:"id,string"`
	GuildId            uint64          `json:"guild_id,string"`
	ChannelId          *uint64         `json:"channel_id,string,omitempty"`
	CreatorId          *uint64         `json:"creator_id,string,omitempty"`
	Name               string          `json:"name"`
	Description        *string         `json:"description,omitempty"`
	ScheduledStartTime time.Time       `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time      `json:"scheduled_end_time,omitempty"`
	PrivacyLevel       PrivacyLevel    `json:"privacy_level"`
	Status             Status          `json:"status"`
	EntityType         EntityType      `json:"entity_type"`
	EntityId           *uint64         `json:"entity_id,string,omitempty"`
	EntityMetadata     *EntityMetadata `json:"entity_metadata,omitempty"`
	Creator            *user.User      `json:"creator,omitempty"`
	UserCount          *int            `json:"user_count,omitempty"`
	Image              *string         `json:"image,omitempty"`
	RecurrenceRule     *RecurrenceRule `json:"recurrence_rule,omitempty"`
}
