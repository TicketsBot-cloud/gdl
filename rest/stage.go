package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/stage"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type CreateStageInstanceData struct {
	ChannelId             uint64                   `json:"channel_id,string"`
	Topic                 string                   `json:"topic"`
	PrivacyLevel          *stage.StagePrivacyLevel `json:"privacy_level,omitempty"`
	SendStartNotification *bool                    `json:"send_start_notification,omitempty"`
	GuildScheduledEventId *uint64                  `json:"guild_scheduled_event_id,string,omitempty"`
}

type ModifyStageInstanceData struct {
	Topic        *string                  `json:"topic,omitempty"`
	PrivacyLevel *stage.StagePrivacyLevel `json:"privacy_level,omitempty"`
}

func CreateStageInstance(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, data CreateStageInstanceData) (stage.StageInstance, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    "/stage-instances",
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteCreateStageInstance, data.ChannelId),
		RateLimiter: rateLimiter,
	}

	var instance stage.StageInstance
	if err, _ := endpoint.Request(ctx, token, data, &instance); err != nil {
		return stage.StageInstance{}, err
	}

	return instance, nil
}

func GetStageInstance(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId uint64) (stage.StageInstance, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/stage-instances/%d", channelId),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteGetStageInstance, channelId),
		RateLimiter: rateLimiter,
	}

	var instance stage.StageInstance
	err, _ := endpoint.Request(ctx, token, nil, &instance)
	return instance, err
}

func ModifyStageInstance(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId uint64, data ModifyStageInstanceData) (stage.StageInstance, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/stage-instances/%d", channelId),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteModifyStageInstance, channelId),
		RateLimiter: rateLimiter,
	}

	var instance stage.StageInstance
	if err, _ := endpoint.Request(ctx, token, data, &instance); err != nil {
		return stage.StageInstance{}, err
	}

	return instance, nil
}

func DeleteStageInstance(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/stage-instances/%d", channelId),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteDeleteStageInstance, channelId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
