package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type GuildMembersChunk struct {
	GuildId   uint64                  `json:"guild_id,string"`
	Members   []member.Member         `json:"member"`
	NotFound  utils.Uint64StringSlice `json:"not_found"`
	Presences []user.Presence         `json:"presences"`
}
