package sticker

import "github.com/TicketsBot-cloud/gdl/objects/user"

type StickerType int

const (
	StickerTypeStandard StickerType = iota + 1
	StickerTypeGuild
)

type StickerFormatType int

const (
	StickerFormatTypePNG    StickerFormatType = iota + 1
	StickerFormatTypeAPNG
	StickerFormatTypeLottie
	StickerFormatTypeGIF
)

type Sticker struct {
	Id          uint64            `json:"id,string"`
	PackId      *uint64           `json:"pack_id,string"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Tags        string            `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   *bool             `json:"available"`
	GuildId     *uint64           `json:"guild_id,string"`
	User        *user.User        `json:"user"`
	SortValue   *int              `json:"sort_value"`
}
