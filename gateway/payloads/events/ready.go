package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type ReadyApplication struct {
	Id    uint64 `json:"id,string"`
	Flags int    `json:"flags"`
}

type Ready struct {
	GatewayVersion   int              `json:"v"`
	User             user.User        `json:"user"`
	Guilds           []guild.Guild    `json:"guilds"`
	SessionId        string           `json:"session_id"`
	ResumeGatewayUrl string           `json:"resume_gateway_url"`
	Shard            []int            `json:"shard"`
	Application      ReadyApplication `json:"application"`
}
