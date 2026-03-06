package guild

import (
	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Webhook struct {
	Id            uint64                    `json:"id,string"`
	Type          WebhookType               `json:"type"`
	GuildId       *uint64                   `json:"guild_id,string,omitempty"`
	ChannelId     objects.NullableSnowflake `json:"channel_id"`
	User          *user.User                `json:"user,omitempty"`
	Name          *string                   `json:"name"`
	Avatar        *string                   `json:"avatar"`
	Token         string                    `json:"token,omitempty"`
	ApplicationId objects.NullableSnowflake `json:"application_id"`
	SourceGuild   *WebhookSourceGuild       `json:"source_guild,omitempty"`
	SourceChannel *channel.PartialChannel   `json:"source_channel,omitempty"`
	Url           *string                   `json:"url,omitempty"`
}

// WebhookSourceGuild is the partial guild object returned for Channel Follower webhooks.
type WebhookSourceGuild struct {
	Id   uint64  `json:"id,string"`
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}
