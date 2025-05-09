package interaction

import (
	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/channel/message"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type ResolvedData struct {
	Users       map[objects.Snowflake]user.User          `json:"users"`
	Members     map[objects.Snowflake]member.Member      `json:"members"`
	Roles       map[objects.Snowflake]guild.Role         `json:"roles"`
	Channels    map[objects.Snowflake]channel.Channel    `json:"channels"`
	Messages    map[objects.Snowflake]message.Message    `json:"messages"`
	Attachments map[objects.Snowflake]channel.Attachment `json:"attachments"`
}
