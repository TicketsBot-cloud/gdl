package rest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/TicketsBot-cloud/gdl/objects/subscription"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type ListSKUSubscriptionsData struct {
	Before *uint64
	After  *uint64
	Limit  *int
	UserId *uint64
}

func (d *ListSKUSubscriptionsData) Query() string {
	query := url.Values{}

	if d.Before != nil {
		query.Add("before", strconv.FormatUint(*d.Before, 10))
	}

	if d.After != nil {
		query.Add("after", strconv.FormatUint(*d.After, 10))
	}

	if d.Limit != nil {
		query.Add("limit", strconv.Itoa(*d.Limit))
	}

	if d.UserId != nil {
		query.Add("user_id", strconv.FormatUint(*d.UserId, 10))
	}

	return query.Encode()
}

func ListSKUSubscriptions(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, skuId uint64, data ListSKUSubscriptionsData) ([]subscription.Subscription, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/skus/%d/subscriptions?%s", skuId, data.Query()),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteListSKUSubscriptions, skuId),
		RateLimiter: rateLimiter,
	}

	var subscriptions []subscription.Subscription
	if err, _ := endpoint.Request(ctx, token, nil, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func GetSKUSubscription(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, skuId, subscriptionId uint64) (subscription.Subscription, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/skus/%d/subscriptions/%d", skuId, subscriptionId),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetSKUSubscription, skuId),
		RateLimiter: rateLimiter,
	}

	var sub subscription.Subscription
	err, _ := endpoint.Request(ctx, token, nil, &sub)
	return sub, err
}
