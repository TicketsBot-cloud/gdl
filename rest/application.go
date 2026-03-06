package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/application"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func GetCurrentApplication(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter) (application.Application, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    "/applications/@me",
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteGetCurrentApplication, 0),
		RateLimiter: rateLimiter,
	}

	var app application.Application
	err, _ := endpoint.Request(ctx, token, nil, &app)
	return app, err
}

type EditCurrentApplicationData struct {
	CustomInstallUrl               *string                                                                                 `json:"custom_install_url,omitempty"`
	Description                    *string                                                                                 `json:"description,omitempty"`
	RoleConnectionsVerificationUrl *string                                                                                 `json:"role_connections_verification_url,omitempty"`
	InstallParams                  *application.InstallParams                                                              `json:"install_params,omitempty"`
	IntegrationTypesConfig         map[application.ApplicationIntegrationType]application.ApplicationIntegrationTypeConfig `json:"integration_types_config,omitempty"`
	Flags                          *application.Flag                                                                       `json:"flags,omitempty"`
	Icon                           *string                                                                                 `json:"icon,omitempty"`
	CoverImage                     *string                                                                                 `json:"cover_image,omitempty"`
	InteractionsEndpointUrl        *string                                                                                 `json:"interactions_endpoint_url,omitempty"`
	Tags                           []string                                                                                `json:"tags,omitempty"`
	EventWebhooksUrl               *string                                                                                 `json:"event_webhooks_url,omitempty"`
	EventWebhooksStatus            *application.ApplicationEventWebhookStatus                                              `json:"event_webhooks_status,omitempty"`
	EventWebhooksTypes             []string                                                                                `json:"event_webhooks_types,omitempty"`
}

func EditCurrentApplication(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, data EditCurrentApplicationData) (application.Application, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    "/applications/@me",
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteEditCurrentApplication, 0),
		RateLimiter: rateLimiter,
	}

	var app application.Application
	err, _ := endpoint.Request(ctx, token, data, &app)
	return app, err
}

func GetActivityInstance(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64, instanceId string) (application.ActivityInstance, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/activity-instances/%s", applicationId, instanceId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteGetActivityInstance, applicationId),
		RateLimiter: rateLimiter,
	}

	var instance application.ActivityInstance
	err, _ := endpoint.Request(ctx, token, nil, &instance)
	return instance, err
}
