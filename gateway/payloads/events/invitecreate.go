package events

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type InviteCreate struct {
	ChannelId  uint64     `json:"channel_id,string"`
	Code       string     `json:"code"`
	CreatedAt  time.Time  `json:"created_at"`
	GuildId    *uint64    `json:"guild_id,string"`
	Inviter    *user.User `json:"inviter"`
	MaxAge     int        `json:"max_age"`
	MaxUses    int        `json:"max_uses"`
	TargetType *int       `json:"target_type"`
	TargetUser *user.User `json:"target_user"`
	Temporary  bool       `json:"temporary"`
	Uses       int        `json:"uses"`
	ExpiresAt  *time.Time `json:"expires_at"`
}
