package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild"
)

type VoiceStateUpdate struct {
	guild.VoiceState
}
