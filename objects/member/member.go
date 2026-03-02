package member

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type Member struct {
	User                       *user.User              `json:"user"`
	Nick                       *string                 `json:"nick"`
	Avatar                     *string                 `json:"avatar"`
	Banner                     *string                 `json:"banner"`
	Roles                      utils.Uint64StringSlice `json:"roles"`
	JoinedAt                   *time.Time              `json:"joined_at"`
	PremiumSince               *time.Time              `json:"premium_since"`
	Deaf                       bool                    `json:"deaf"`
	Mute                       bool                    `json:"mute"`
	Flags                      int                      `json:"flags"`
	Pending                    *bool                    `json:"pending,omitempty"`
	Permissions                uint64                   `json:"permissions,string,omitempty"`
	CommunicationDisabledUntil *time.Time               `json:"communication_disabled_until,omitempty"`
	AvatarDecorationData       *user.AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
}

func (m *Member) HasRole(roleId uint64) bool {
	for _, memberRole := range m.Roles {
		if memberRole == roleId {
			return true
		}
	}
	return false
}

func (m *Member) ToCachedMember() CachedMember {
	var joinedAt time.Time
	if m.JoinedAt != nil {
		joinedAt = *m.JoinedAt
	}

	return CachedMember{
		Nick:         m.Nick,
		Roles:        m.Roles,
		JoinedAt:     joinedAt,
		PremiumSince: m.PremiumSince,
		Deaf:         m.Deaf,
		Mute:         m.Mute,
	}
}
