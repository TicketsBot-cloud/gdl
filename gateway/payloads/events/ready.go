package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/application"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Ready struct {
	GatewayVersion   int                            `json:"v"`
	User             user.User                      `json:"user"`
	Guilds           []guild.Guild                  `json:"guilds"`
	SessionId        string                         `json:"session_id"`
	ResumeGatewayUrl string                         `json:"resume_gateway_url"`
	Shard            []int                          `json:"shard"`
	Application      application.PartialApplication `json:"application"`
}
