package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/automoderation"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type CreateAutoModerationRuleData struct {
	Name            string                          `json:"name"`
	EventType       automoderation.EventType        `json:"event_type"`
	TriggerType     automoderation.TriggerType      `json:"trigger_type"`
	TriggerMetadata *automoderation.TriggerMetadata `json:"trigger_metadata,omitempty"`
	Actions         []automoderation.Action         `json:"actions"`
	Enabled         *bool                           `json:"enabled,omitempty"`
	ExemptRoles     utils.Uint64StringSlice         `json:"exempt_roles,omitempty"`
	ExemptChannels  utils.Uint64StringSlice         `json:"exempt_channels,omitempty"`
}

type ModifyAutoModerationRuleData struct {
	Name            *string                         `json:"name,omitempty"`
	EventType       *automoderation.EventType       `json:"event_type,omitempty"`
	TriggerMetadata *automoderation.TriggerMetadata `json:"trigger_metadata,omitempty"`
	Actions         []automoderation.Action         `json:"actions,omitempty"`
	Enabled         *bool                           `json:"enabled,omitempty"`
	ExemptRoles     utils.Uint64StringSlice         `json:"exempt_roles,omitempty"`
	ExemptChannels  utils.Uint64StringSlice         `json:"exempt_channels,omitempty"`
}

func ListAutoModerationRules(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) ([]automoderation.Rule, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/auto-moderation/rules", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteListAutoModerationRules, guildId),
		RateLimiter: rateLimiter,
	}

	var rules []automoderation.Rule
	if err, _ := endpoint.Request(ctx, token, nil, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

func GetAutoModerationRule(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, ruleId uint64) (automoderation.Rule, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/auto-moderation/rules/%d", guildId, ruleId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetAutoModerationRule, guildId),
		RateLimiter: rateLimiter,
	}

	var rule automoderation.Rule
	err, _ := endpoint.Request(ctx, token, nil, &rule)
	return rule, err
}

func CreateAutoModerationRule(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, data CreateAutoModerationRuleData) (automoderation.Rule, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/auto-moderation/rules", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteCreateAutoModerationRule, guildId),
		RateLimiter: rateLimiter,
	}

	var rule automoderation.Rule
	if err, _ := endpoint.Request(ctx, token, data, &rule); err != nil {
		return automoderation.Rule{}, err
	}

	return rule, nil
}

func ModifyAutoModerationRule(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, ruleId uint64, data ModifyAutoModerationRuleData) (automoderation.Rule, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/auto-moderation/rules/%d", guildId, ruleId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyAutoModerationRule, guildId),
		RateLimiter: rateLimiter,
	}

	var rule automoderation.Rule
	if err, _ := endpoint.Request(ctx, token, data, &rule); err != nil {
		return automoderation.Rule{}, err
	}

	return rule, nil
}

func DeleteAutoModerationRule(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, ruleId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/auto-moderation/rules/%d", guildId, ruleId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteAutoModerationRule, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
