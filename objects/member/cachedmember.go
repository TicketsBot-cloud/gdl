package member

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type CachedMember struct {
	Nick         string     `json:"nick"`
	Roles        []uint64   `json:"roles"`
	JoinedAt     time.Time  `json:"joined_at"`
	PremiumSince *time.Time `json:"premium_since"` // when the user started boosting the guild
	Deaf         bool       `json:"deaf"`
	Mute         bool       `json:"mute"`
}

func (m *CachedMember) ToMember(user user.User) Member {
	return Member{
		User:         user,
		Nick:         m.Nick,
		Roles:        m.Roles,
		JoinedAt:     m.JoinedAt,
		PremiumSince: m.PremiumSince,
		Deaf:         m.Deaf,
		Mute:         m.Mute,
	}
}
