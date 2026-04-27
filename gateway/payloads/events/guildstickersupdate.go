package events

import "github.com/TicketsBot-cloud/gdl/objects/guild/sticker"

type GuildStickersUpdate struct {
	GuildId  uint64            `json:"guild_id,string"`
	Stickers []sticker.Sticker `json:"stickers"`
}
