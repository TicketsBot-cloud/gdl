package events

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type ThreadCreate struct {
	channel.Channel
	NewlyCreated bool `json:"newly_created"`
}
