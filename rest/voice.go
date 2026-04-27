package rest

import (
	"context"
	"fmt"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func ListVoiceRegions(ctx context.Context, token string) ([]guild.VoiceRegion, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    "/voice/regions",
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteListVoiceRegions, 0),
	}

	var voiceRegions []guild.VoiceRegion
	err, _ := endpoint.Request(ctx, token, nil, &voiceRegions)
	return voiceRegions, err
}

func GetCurrentUserVoiceState(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) (guild.VoiceState, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/voice-states/@me", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetCurrentUserVoiceState, guildId),
		RateLimiter: rateLimiter,
	}

	var state guild.VoiceState
	err, _ := endpoint.Request(ctx, token, nil, &state)
	return state, err
}

func GetUserVoiceState(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, userId uint64) (guild.VoiceState, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/voice-states/%d", guildId, userId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetUserVoiceState, guildId),
		RateLimiter: rateLimiter,
	}

	var state guild.VoiceState
	err, _ := endpoint.Request(ctx, token, nil, &state)
	return state, err
}

type ModifyCurrentUserVoiceStateData struct {
	ChannelId               *uint64    `json:"channel_id,string,omitempty"`
	Suppress                *bool      `json:"suppress,omitempty"`
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
}

func ModifyCurrentUserVoiceState(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, data ModifyCurrentUserVoiceStateData) error {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/voice-states/@me", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyCurrentUserVoiceState, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, data, nil)
	return err
}

type ModifyUserVoiceStateData struct {
	ChannelId *uint64 `json:"channel_id,string,omitempty"`
	Suppress  *bool   `json:"suppress,omitempty"`
}

func ModifyUserVoiceState(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, userId uint64, data ModifyUserVoiceStateData) error {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/voice-states/%d", guildId, userId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyUserVoiceState, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, data, nil)
	return err
}
