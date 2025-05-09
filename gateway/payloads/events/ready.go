package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Ready struct {
	GatewayVersion  int           `json:"v"`
	User            user.User     `json:"user"`
	PrivateChannels []uint64      `json:"private_channels,string"` // Note: This slice will always be empty
	Guilds          []guild.Guild `json:"guilds"`
	SessionId       string        `json:"session_id"`
	Shard           []int         `json:"shard"` // Slice of [shard_id, num_shards]
}
