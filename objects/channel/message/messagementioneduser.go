package message

import (
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

// Mentions is an array of users with partial member
type MessageMentionedUser struct {
	user.User
	Member member.Member
}
