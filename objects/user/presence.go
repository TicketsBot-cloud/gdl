package user

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/utils"
)

type Presence struct {
	User         User                    `json:"user"`
	GuildId      *uint64                 `json:"guild_id,string,omitempty"`
	Status       string                  `json:"status"`
	Activities   []Activity              `json:"activities"`
	ClientStatus ClientStatus            `json:"client_status"`
	Roles        utils.Uint64StringSlice `json:"roles,omitempty"`
	PremiumSince *time.Time              `json:"premium_since,omitempty"`
	Nick         *string                 `json:"nick,omitempty"`
}
