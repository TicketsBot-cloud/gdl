package emoji

import (
	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type CachedEmoji struct {
	GuildId       uint64   `json:"-"`
	Name          string   `json:"name"`
	Roles         []uint64 `json:"roles"`
	User          uint64   `json:"user"`
	RequireColons bool     `json:"require_colons"`
	Managed       bool     `json:"managed"`
	Animated      bool     `json:"animated"`
}

func (e *CachedEmoji) ToEmoji(emojiId uint64, user user.User) Emoji {
	return Emoji{
		Id:            objects.NewNullableSnowflake(emojiId),
		Name:          e.Name,
		Roles:         e.Roles,
		User:          user,
		RequireColons: e.RequireColons,
		Managed:       e.Managed,
		Animated:      e.Animated,
	}
}
