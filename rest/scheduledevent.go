package rest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/guild/scheduledevent"
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type CreateGuildScheduledEventData struct {
	ChannelId          *uint64                        `json:"channel_id,string,omitempty"`
	EntityMetadata     *scheduledevent.EntityMetadata `json:"entity_metadata,omitempty"`
	Name               string                         `json:"name"`
	PrivacyLevel       scheduledevent.PrivacyLevel    `json:"privacy_level"`
	ScheduledStartTime time.Time                      `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                     `json:"scheduled_end_time,omitempty"`
	Description        *string                        `json:"description,omitempty"`
	EntityType         scheduledevent.EntityType      `json:"entity_type"`
	Image              *string                        `json:"image,omitempty"`
	RecurrenceRule     *scheduledevent.RecurrenceRule `json:"recurrence_rule,omitempty"`
}

type ModifyGuildScheduledEventData struct {
	ChannelId          *uint64                         `json:"channel_id,string,omitempty"`
	EntityMetadata     *scheduledevent.EntityMetadata  `json:"entity_metadata,omitempty"`
	Name               *string                         `json:"name,omitempty"`
	PrivacyLevel       *scheduledevent.PrivacyLevel    `json:"privacy_level,omitempty"`
	ScheduledStartTime *time.Time                      `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime   *time.Time                      `json:"scheduled_end_time,omitempty"`
	Description        *string                         `json:"description,omitempty"`
	EntityType         *scheduledevent.EntityType      `json:"entity_type,omitempty"`
	Status             *scheduledevent.Status          `json:"status,omitempty"`
	Image              *string                         `json:"image,omitempty"`
	RecurrenceRule     *scheduledevent.RecurrenceRule  `json:"recurrence_rule,omitempty"`
}

type GetGuildScheduledEventUsersData struct {
	Limit      *int
	WithMember *bool
	Before     *uint64
	After      *uint64
}

func (d *GetGuildScheduledEventUsersData) Query() string {
	query := url.Values{}

	if d.Limit != nil {
		query.Add("limit", strconv.Itoa(*d.Limit))
	}

	if d.WithMember != nil {
		query.Add("with_member", strconv.FormatBool(*d.WithMember))
	}

	if d.Before != nil {
		query.Add("before", strconv.FormatUint(*d.Before, 10))
	}

	if d.After != nil {
		query.Add("after", strconv.FormatUint(*d.After, 10))
	}

	return query.Encode()
}

type GuildScheduledEventUser struct {
	GuildScheduledEventId uint64         `json:"guild_scheduled_event_id,string"`
	User                  user.User      `json:"user"`
	Member                *member.Member `json:"member,omitempty"`
}

func ListGuildScheduledEvents(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, withUserCount bool) ([]scheduledevent.GuildScheduledEvent, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events?with_user_count=%s", guildId, strconv.FormatBool(withUserCount)),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteListGuildScheduledEvents, guildId),
		RateLimiter: rateLimiter,
	}

	var events []scheduledevent.GuildScheduledEvent
	if err, _ := endpoint.Request(ctx, token, nil, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func GetGuildScheduledEvent(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, eventId uint64, withUserCount bool) (scheduledevent.GuildScheduledEvent, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events/%d?with_user_count=%s", guildId, eventId, strconv.FormatBool(withUserCount)),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetGuildScheduledEvent, guildId),
		RateLimiter: rateLimiter,
	}

	var event scheduledevent.GuildScheduledEvent
	err, _ := endpoint.Request(ctx, token, nil, &event)
	return event, err
}

func CreateGuildScheduledEvent(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, data CreateGuildScheduledEventData) (scheduledevent.GuildScheduledEvent, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteCreateGuildScheduledEvent, guildId),
		RateLimiter: rateLimiter,
	}

	var event scheduledevent.GuildScheduledEvent
	if err, _ := endpoint.Request(ctx, token, data, &event); err != nil {
		return scheduledevent.GuildScheduledEvent{}, err
	}

	return event, nil
}

func ModifyGuildScheduledEvent(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, eventId uint64, data ModifyGuildScheduledEventData) (scheduledevent.GuildScheduledEvent, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events/%d", guildId, eventId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyGuildScheduledEvent, guildId),
		RateLimiter: rateLimiter,
	}

	var event scheduledevent.GuildScheduledEvent
	if err, _ := endpoint.Request(ctx, token, data, &event); err != nil {
		return scheduledevent.GuildScheduledEvent{}, err
	}

	return event, nil
}

func DeleteGuildScheduledEvent(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, eventId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events/%d", guildId, eventId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteGuildScheduledEvent, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}

func GetGuildScheduledEventUsers(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, eventId uint64, data GetGuildScheduledEventUsersData) ([]GuildScheduledEventUser, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/scheduled-events/%d/users?%s", guildId, eventId, data.Query()),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetGuildScheduledEventUsers, guildId),
		RateLimiter: rateLimiter,
	}

	var users []GuildScheduledEventUser
	if err, _ := endpoint.Request(ctx, token, nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}
