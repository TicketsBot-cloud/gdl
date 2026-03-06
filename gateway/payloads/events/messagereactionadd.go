package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
	"github.com/TicketsBot-cloud/gdl/objects/member"
)

type MessageReactionAdd struct {
	UserId          uint64         `json:"user_id,string"`
	ChannelId       uint64         `json:"channel_id,string"`
	MessageId       uint64         `json:"message_id,string"`
	GuildId         *uint64        `json:"guild_id,string"`
	MessageAuthorId *uint64        `json:"message_author_id,string"`
	Member          *member.Member `json:"member"`
	Emoji           emoji.Emoji    `json:"emoji"`
	Burst           bool           `json:"burst"`
	BurstColors     []string       `json:"burst_colors"`
	Type            int            `json:"type"`
}
