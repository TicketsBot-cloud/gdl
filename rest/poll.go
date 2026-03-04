package rest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/TicketsBot-cloud/gdl/objects/channel/message"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type GetPollAnswerVotersData struct {
	After *uint64
	Limit *int
}

func (d *GetPollAnswerVotersData) Query() string {
	query := url.Values{}

	if d.After != nil {
		query.Add("after", strconv.FormatUint(*d.After, 10))
	}

	if d.Limit != nil {
		query.Add("limit", strconv.Itoa(*d.Limit))
	}

	return query.Encode()
}

type pollAnswerVotersResponse struct {
	Users []user.User `json:"users"`
}

func GetPollAnswerVoters(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId, messageId uint64, answerId int, data GetPollAnswerVotersData) ([]user.User, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/channels/%d/polls/%d/answers/%d?%s", channelId, messageId, answerId, data.Query()),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteGetPollAnswerVoters, channelId),
		RateLimiter: rateLimiter,
	}

	var res pollAnswerVotersResponse
	if err, _ := endpoint.Request(ctx, token, nil, &res); err != nil {
		return nil, err
	}

	return res.Users, nil
}

func EndPoll(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId, messageId uint64) (message.Message, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/channels/%d/polls/%d/expire", channelId, messageId),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteEndPoll, channelId),
		RateLimiter: rateLimiter,
	}

	var msg message.Message
	if err, _ := endpoint.Request(ctx, token, nil, &msg); err != nil {
		return message.Message{}, err
	}

	return msg, nil
}
