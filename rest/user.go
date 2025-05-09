package rest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/integration"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func GetCurrentUser(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter) (user.User, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    "/users/@me",
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetCurrentUser, 0),
		RateLimiter: rateLimiter,
	}

	var user user.User
	err, _ := endpoint.Request(ctx, token, nil, &user)
	return user, err
}

func GetUser(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, userId uint64) (user.User, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/users/%d", userId),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetUser, userId),
		RateLimiter: rateLimiter,
	}

	var user user.User
	err, _ := endpoint.Request(ctx, token, nil, &user)
	return user, err
}

type ModifyUserData struct {
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

func ModifyCurrentUser(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, data ModifyUserData) (user.User, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    "/users/@me",
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteModifyCurrentUser, 0),
		RateLimiter: rateLimiter,
	}

	var user user.User
	err, _ := endpoint.Request(ctx, token, data, &user)
	return user, err
}

type CurrentUserGuildsData struct {
	Before uint64
	After  uint64
	Limit  int
}

func (d *CurrentUserGuildsData) Query() string {
	query := url.Values{}

	if d.Before != 0 {
		query.Set("before", strconv.FormatUint(d.Before, 10))
	}

	if d.After != 0 {
		query.Set("after", strconv.FormatUint(d.After, 10))
	}

	if d.Limit > 200 || d.Limit < 1 {
		d.Limit = 200
	}
	query.Set("limit", strconv.Itoa(d.Limit))

	return query.Encode()
}

func GetCurrentUserGuilds(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, data CurrentUserGuildsData) ([]guild.Guild, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/users/@me/guilds?%s", data.Query()),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetCurrentUserGuilds, 0),
		RateLimiter: rateLimiter,
	}

	var guilds []guild.Guild
	err, _ := endpoint.Request(ctx, token, nil, &guilds)
	return guilds, err
}

func LeaveGuild(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/users/@me/guilds/%d", guildId),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteLeaveGuild, 0),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}

func CreateDM(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, recipientId uint64) (channel.Channel, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/users/@me/channels"),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteCreateDM, recipientId),
		RateLimiter: rateLimiter,
	}

	body := map[string]interface{}{
		"recipient_id": strconv.FormatUint(recipientId, 10),
	}

	var channel channel.Channel
	err, _ := endpoint.Request(ctx, token, body, &channel)
	return channel, err
}

func GetUserConnections(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter) ([]integration.Connection, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/users/@me/connections"),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetUserConnections, 0),
		RateLimiter: rateLimiter,
	}

	var connections []integration.Connection
	err, _ := endpoint.Request(ctx, token, nil, &connections)
	return connections, err
}
