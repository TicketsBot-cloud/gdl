package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/sku"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func ListSKUs(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64) ([]sku.SKU, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/skus", applicationId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteListSKUs, applicationId),
		RateLimiter: rateLimiter,
	}

	var skus []sku.SKU
	if err, _ := endpoint.Request(ctx, token, nil, &skus); err != nil {
		return nil, err
	}

	return skus, nil
}
