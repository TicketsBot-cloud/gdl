package member

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type CachedMember struct {
	Nick         *string    `json:"nick"`
	Roles        []uint64   `json:"roles"`
	JoinedAt     time.Time  `json:"joined_at"`
	PremiumSince *time.Time `json:"premium_since"`
	Deaf         bool       `json:"deaf"`
	Mute         bool       `json:"mute"`
}

func (m *CachedMember) ToMember(u user.User) Member {
	return Member{
		User:         &u,
		Nick:         m.Nick,
		Roles:        m.Roles,
		JoinedAt:     &m.JoinedAt,
		PremiumSince: m.PremiumSince,
		Deaf:         m.Deaf,
		Mute:         m.Mute,
	}
}
