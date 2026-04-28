package events

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type ChannelInfo struct {
	GuildId  uint64                     `json:"guild_id,string"`
	Channels []channel.ChannelInfoEntry `json:"channels"`
}
