package command

import (
	"github.com/TicketsBot-cloud/gdl/gateway"
	"github.com/TicketsBot-cloud/gdl/gateway/payloads/events"
)

type CommandContext struct {
	*events.MessageCreate
	Shard *gateway.Shard
	Args  []string
}
