package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type GuildBanRemove struct {
	GuildId uint64    `json:"guild_id,string"`
	User    user.User `json:"user"`
}
