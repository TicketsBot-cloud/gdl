package user

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/utils"
)

type Presence struct {
	User         User                    `json:"user"`
	Roles        utils.Uint64StringSlice `json:"roles"`
	Game         Activity                `json:"name"`
	GuildId      uint64                  `json:"guild_id,string"`
	Status       string                  `json:"status"`
	Activities   []Activity              `json:"activities"`
	ClientStatus ClientStatus            `json:"client_status"`
	PremiumSince *time.Time              `json:"premium_since"`
	Nick         string                  `json:"nick"`
}
