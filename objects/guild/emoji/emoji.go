package emoji

import (
	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

// https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	Id            objects.NullableSnowflake `json:"id"`
	Name          *string                   `json:"name,omitempty"` // null when custom emoji data is not available; unicode emoji string for standard emoji
	Roles         utils.Uint64StringSlice   `json:"roles,omitempty"`
	User          *user.User                `json:"user,omitempty"`
	RequireColons bool                      `json:"require_colons,omitempty"`
	Managed       bool                      `json:"managed,omitempty"`
	Animated      bool                      `json:"animated,omitempty"`
	Available     *bool                     `json:"available,omitempty"`
}

func (e *Emoji) ToCachedEmoji(guildId uint64) CachedEmoji {
	var userId uint64
	if e.User != nil {
		userId = e.User.Id
	}

	return CachedEmoji{
		GuildId:       guildId,
		Name:          e.Name,
		Roles:         e.Roles,
		User:          userId,
		RequireColons: e.RequireColons,
		Managed:       e.Managed,
		Animated:      e.Animated,
	}
}
