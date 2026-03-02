package events

import "github.com/TicketsBot-cloud/gdl/objects/auditlog"

type GuildAuditLogEntryCreate struct {
	auditlog.AuditLogEntry
	GuildId uint64 `json:"guild_id,string"`
}
