package auditlog

import (
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/integration"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type AuditLog struct {
	Webhooks     []guild.Webhook           `json:"webhooks"`
	Users        []user.User               `json:"users"`
	Entries      []AuditLogEntry           `json:"audit_log_entries"`
	Integrations []integration.Integration `json:"integrations"`
}
