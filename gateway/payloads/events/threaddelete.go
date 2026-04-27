package events

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type ThreadDelete struct {
	Id       uint64              `json:"id,string"`
	GuildId  uint64              `json:"guild_id,string"`
	ParentId uint64              `json:"parent_id,string"`
	Type     channel.ChannelType `json:"type"`
}
