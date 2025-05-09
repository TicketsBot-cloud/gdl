package events

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type GuildMemberUpdate struct {
	GuildId      uint64                  `json:"guild_id,string"`
	Roles        utils.Uint64StringSlice `json:"roles"`
	User         user.User               `json:"user"`
	Nick         string                  `json:"nick"`
	PremiumSince *time.Time              `json:"premium_since"` // When the user started boosting the guild
}
