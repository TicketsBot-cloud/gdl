package message

import "github.com/TicketsBot-cloud/gdl/objects/guild/sticker"

type StickerItem struct {
	Id         uint64                    `json:"id,string"`
	Name       string                    `json:"name"`
	FormatType sticker.StickerFormatType `json:"format_type"`
}
