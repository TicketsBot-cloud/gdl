package events

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type ThreadMemberUpdate struct {
	channel.ThreadMember
	GuildId uint64 `json:"guild_id,string"`
}
