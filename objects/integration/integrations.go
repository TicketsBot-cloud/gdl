package integration

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Integration struct {
	Id                uint64                     `json:"id,string"`
	Name              string                     `json:"name"`
	Type              string                     `json:"type"` // twitch, youtube, etc.
	Enabled           bool                       `json:"enabled"`
	Syncing           bool                       `json:"syncing"`
	RoleId            uint64                     `json:"role_id,string"`
	EnableEmoticons   bool                       `json:"enable_emoticons"`
	ExpireBehaviour   IntegrationExpireBehaviour `json:"expire_behavior"`
	ExpireGracePeriod int                        `json:"expire_grace_period"`
	User              user.User                  `json:"user"`
	Account           guild.Account              `json:"account"`
	SyncedAt          time.Time                  `json:"synced_at"`
}
