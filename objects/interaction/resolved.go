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
	Users       map[objects.Snowflake]user.User          `json:"users,omitempty"`
	Members     map[objects.Snowflake]member.Member      `json:"members,omitempty"`
	Roles       map[objects.Snowflake]guild.Role         `json:"roles,omitempty"`
	Channels    map[objects.Snowflake]channel.Channel    `json:"channels,omitempty"`
	Messages    map[objects.Snowflake]message.Message    `json:"messages,omitempty"`
	Attachments map[objects.Snowflake]channel.Attachment `json:"attachments,omitempty"`
}
