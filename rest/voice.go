package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

func ListVoiceRegions(ctx context.Context, token string) ([]guild.VoiceRegion, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/voice/regions"),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteListVoiceRegions, 0),
	}

	var voiceRegions []guild.VoiceRegion
	err, _ := endpoint.Request(ctx, token, nil, &voiceRegions)
	return voiceRegions, err
}
