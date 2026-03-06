package sticker

import "github.com/TicketsBot-cloud/gdl/objects/user"

type StickerType int

const (
	StickerTypeStandard StickerType = iota + 1
	StickerTypeGuild
)

type StickerFormatType int

const (
	StickerFormatTypePNG StickerFormatType = iota + 1
	StickerFormatTypeAPNG
	StickerFormatTypeLottie
	StickerFormatTypeGIF
)

type Sticker struct {
	Id          uint64            `json:"id,string"`
	PackId      *uint64           `json:"pack_id,string,omitempty"`
	Name        string            `json:"name"`
	Description *string           `json:"description,omitempty"`
	Tags        string            `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   *bool             `json:"available,omitempty"`
	GuildId     *uint64           `json:"guild_id,string,omitempty"`
	User        *user.User        `json:"user,omitempty"`
	SortValue   *int              `json:"sort_value,omitempty"`
}
