package events

import "github.com/TicketsBot-cloud/gdl/objects/guild/scheduledevent"

type GuildScheduledEventCreate struct {
	scheduledevent.GuildScheduledEvent
}

type GuildScheduledEventUpdate struct {
	scheduledevent.GuildScheduledEvent
}

type GuildScheduledEventDelete struct {
	scheduledevent.GuildScheduledEvent
}

type GuildScheduledEventUserAdd struct {
	GuildScheduledEventId uint64 `json:"guild_scheduled_event_id,string"`
	UserId                uint64 `json:"user_id,string"`
	GuildId               uint64 `json:"guild_id,string"`
}

type GuildScheduledEventUserRemove struct {
	GuildScheduledEventId uint64 `json:"guild_scheduled_event_id,string"`
	UserId                uint64 `json:"user_id,string"`
	GuildId               uint64 `json:"guild_id,string"`
}
