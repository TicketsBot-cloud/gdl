package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type ThreadListSync struct {
	GuildId    uint64                  `json:"guild_id,string"`
	ChannelIds utils.Uint64StringSlice `json:"channel_ids"`
	Threads    []channel.Channel       `json:"threads"`
	Members    []channel.ThreadMember  `json:"members"`
}
