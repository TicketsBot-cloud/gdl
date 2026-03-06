package events

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type GuildMemberUpdate struct {
	GuildId                    uint64                  `json:"guild_id,string"`
	Roles                      utils.Uint64StringSlice `json:"roles"`
	User                       user.User               `json:"user"`
	Nick                       *string                 `json:"nick"`
	Avatar                     *string                 `json:"avatar"`
	Banner                     *string                 `json:"banner"`
	JoinedAt                   *time.Time              `json:"joined_at"`
	PremiumSince               *time.Time              `json:"premium_since"`
	Deaf                       *bool                   `json:"deaf"`
	Mute                       *bool                   `json:"mute"`
	Pending                    *bool                   `json:"pending"`
	CommunicationDisabledUntil *time.Time              `json:"communication_disabled_until"`
}
