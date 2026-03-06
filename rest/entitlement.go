package rest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/TicketsBot-cloud/gdl/objects/entitlement"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type EntitlementQueryOptions struct {
	UserId         *uint64
	SkuIds         []uint64
	Before         *uint64
	After          *uint64
	Limit          *int
	GuildId        *uint64
	ExcludeEnded   *bool
	ExcludeDeleted *bool
}

func (o *EntitlementQueryOptions) Query() string {
	query := url.Values{}

	if o.UserId != nil {
		query.Add("user_id", strconv.FormatUint(*o.UserId, 10))
	}

	if len(o.SkuIds) > 0 {
		var encoded string
		for i, id := range o.SkuIds {
			encoded += strconv.FormatUint(id, 10)
			if i < len(o.SkuIds)-1 {
				encoded += ","
			}
		}
		query.Add("sku_ids", encoded)
	}

	if o.Before != nil {
		query.Add("before", strconv.FormatUint(*o.Before, 10))
	}

	if o.After != nil {
		query.Add("after", strconv.FormatUint(*o.After, 10))
	}

	if o.Limit != nil {
		query.Add("limit", strconv.Itoa(*o.Limit))
	}

	if o.GuildId != nil {
		query.Add("guild_id", strconv.FormatUint(*o.GuildId, 10))
	}

	if o.ExcludeEnded != nil {
		query.Add("exclude_ended", strconv.FormatBool(*o.ExcludeEnded))
	}

	if o.ExcludeDeleted != nil {
		query.Add("exclude_deleted", strconv.FormatBool(*o.ExcludeDeleted))
	}

	return query.Encode()
}

func ListEntitlements(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64, options EntitlementQueryOptions) ([]entitlement.Entitlement, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/entitlements?%s", applicationId, options.Query()),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteListEntitlements, applicationId),
		RateLimiter: rateLimiter,
	}

	var entitlements []entitlement.Entitlement
	if err, _ := endpoint.Request(ctx, token, nil, &entitlements); err != nil {
		return nil, err
	}

	return entitlements, nil
}

func GetEntitlement(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, entitlementId uint64) (entitlement.Entitlement, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/entitlements/%d", applicationId, entitlementId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteGetEntitlement, applicationId),
		RateLimiter: rateLimiter,
	}

	var e entitlement.Entitlement
	err, _ := endpoint.Request(ctx, token, nil, &e)
	return e, err
}

func ConsumeEntitlement(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, entitlementId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/entitlements/%d/consume", applicationId, entitlementId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteConsumeEntitlement, applicationId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}

type CreateTestEntitlementData struct {
	SkuId     uint64               `json:"sku_id,string"`
	OwnerId   uint64               `json:"owner_id,string"`
	OwnerType EntitlementOwnerType `json:"owner_type"`
}

type EntitlementOwnerType uint8

const (
	EntitlementOwnerTypeGuild EntitlementOwnerType = iota + 1
	EntitlementOwnerTypeUser
)

func CreateTestEntitlement(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64, data CreateTestEntitlementData) (entitlement.Entitlement, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/applications/%d/entitlements", applicationId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteCreateTestEntitlement, applicationId),
		RateLimiter: rateLimiter,
	}

	var e entitlement.Entitlement
	if err, _ := endpoint.Request(ctx, token, data, &e); err != nil {
		return entitlement.Entitlement{}, err
	}

	return e, nil
}

func DeleteTestEntitlement(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, entitlementId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/entitlements/%d", applicationId, entitlementId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteDeleteTestEntitlement, applicationId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
