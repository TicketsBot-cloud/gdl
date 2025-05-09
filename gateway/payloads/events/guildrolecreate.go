package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild"
)

type GuildRoleCreate struct {
	GuildId uint64     `json:"guild_id,string"`
	Role    guild.Role ` json:"role"`
}
