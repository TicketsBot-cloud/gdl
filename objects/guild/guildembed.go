package guild

import "github.com/TicketsBot-cloud/gdl/objects"

// GuildEmbed represents the settings for a guild's widget (also known as GuildWidgetSettings).
type GuildEmbed struct {
	Enabled   bool                      `json:"enabled"`
	ChannelId objects.NullableSnowflake `json:"channel_id"`
}
