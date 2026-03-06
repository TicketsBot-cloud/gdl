package invite

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/application"
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/guild/scheduledevent"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type InviteType int

const (
	InviteTypeGuild InviteType = iota
	InviteTypeGroupDm
	InviteTypeFriend
)

type InviteTargetType int

const (
	InviteTargetTypeStream InviteTargetType = iota + 1
	InviteTargetTypeEmbeddedApplication
)

type Invite struct {
	Type                     InviteType                          `json:"type"`
	Code                     string                              `json:"code"`
	Guild                    *guild.Guild                        `json:"guild,omitempty"`
	Channel                  *channel.PartialChannel             `json:"channel,omitempty"`
	Inviter                  *user.User                          `json:"inviter,omitempty"`
	TargetType               InviteTargetType                    `json:"target_type,omitempty"`
	TargetUser               *user.User                          `json:"target_user,omitempty"`
	TargetApplication        *application.Application            `json:"target_application,omitempty"`
	ApproximatePresenceCount int                                 `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int                                 `json:"approximate_member_count,omitempty"`
	ExpiresAt                *time.Time                          `json:"expires_at,omitempty"`
	GuildScheduledEvent      *scheduledevent.GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
	Flags                    *int                                `json:"flags,omitempty"`
	Roles                    []guild.Role                        `json:"roles,omitempty"`
}

type InviteMetadata struct {
	Invite
	Uses      int       `json:"uses"`
	MaxUses   int       `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
}
