package guild

import (
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Ban struct {
	Reason string    `json:"reason,omitempty"`
	User   user.User `json:"user"`
}
