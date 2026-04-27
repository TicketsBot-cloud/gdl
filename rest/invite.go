package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/invite"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func GetInvite(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, inviteCode string, withCounts bool, guildScheduledEventId *uint64) (invite.Invite, error) {
	url := fmt.Sprintf("/invites/%s?with_counts=%v", inviteCode, withCounts)
	if guildScheduledEventId != nil {
		url += fmt.Sprintf("&guild_scheduled_event_id=%d", *guildScheduledEventId)
	}

	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    url,
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetInvite, 0), // No ratelimit
		RateLimiter: rateLimiter,
	}

	var invite invite.Invite
	err, _ := endpoint.Request(ctx, token, nil, &invite)
	return invite, err
}

func DeleteInvite(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, inviteCode string) (invite.Invite, error) {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    "/invites/" + inviteCode,
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteInvite, 0), // No ratelimit
		RateLimiter: rateLimiter,
	}

	var invite invite.Invite
	err, _ := endpoint.Request(ctx, token, nil, &invite)
	return invite, err
}
