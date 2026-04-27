package events

import "github.com/TicketsBot-cloud/gdl/objects/integration"

type IntegrationCreate struct {
	integration.Integration
	GuildId uint64 `json:"guild_id,string"`
}

type IntegrationUpdate struct {
	integration.Integration
	GuildId uint64 `json:"guild_id,string"`
}

type IntegrationDelete struct {
	Id            uint64  `json:"id,string"`
	GuildId       uint64  `json:"guild_id,string"`
	ApplicationId *uint64 `json:"application_id,string"`
}
