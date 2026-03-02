package message

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/utils"
)

type MessageCall struct {
	Participants   utils.Uint64StringSlice `json:"participants"`
	EndedTimestamp *time.Time              `json:"ended_timestamp,omitempty"`
}
