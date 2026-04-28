package integration

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Integration struct {
	Id                uint64                      `json:"id,string"`
	Name              string                      `json:"name"`
	Type              string                      `json:"type"` // twitch, youtube, discord, etc.
	Enabled           bool                        `json:"enabled"`
	Syncing           *bool                       `json:"syncing,omitempty"`
	RoleId            *uint64                     `json:"role_id,string,omitempty"`
	EnableEmoticons   *bool                       `json:"enable_emoticons,omitempty"`
	ExpireBehaviour   *IntegrationExpireBehaviour `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *int                        `json:"expire_grace_period,omitempty"`
	User              *user.User                  `json:"user,omitempty"`
	Account           guild.Account               `json:"account"`
	SyncedAt          *time.Time                  `json:"synced_at,omitempty"`
	SubscriberCount   *int                        `json:"subscriber_count,omitempty"`
	Revoked           *bool                       `json:"revoked,omitempty"`
	Application       *IntegrationApplication     `json:"application,omitempty"`
	Scopes            []string                    `json:"scopes,omitempty"`
}

type IntegrationApplication struct {
	Id          uint64     `json:"id,string"`
	Name        string     `json:"name"`
	Icon        *string    `json:"icon"`
	Description string     `json:"description"`
	Bot         *user.User `json:"bot,omitempty"`
}
