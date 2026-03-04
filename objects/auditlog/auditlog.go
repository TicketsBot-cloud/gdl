package auditlog

import (
	"github.com/TicketsBot-cloud/gdl/objects/automoderation"
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/guild/scheduledevent"
	"github.com/TicketsBot-cloud/gdl/objects/integration"
	"github.com/TicketsBot-cloud/gdl/objects/interaction"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type AuditLog struct {
	ApplicationCommands  []interaction.ApplicationCommand     `json:"application_commands"`
	AutoModerationRules  []automoderation.Rule                `json:"auto_moderation_rules"`
	Entries              []AuditLogEntry                      `json:"audit_log_entries"`
	GuildScheduledEvents []scheduledevent.GuildScheduledEvent `json:"guild_scheduled_events"`
	Integrations         []integration.Integration            `json:"integrations"`
	Threads              []channel.Channel                    `json:"threads"`
	Users                []user.User                          `json:"users"`
	Webhooks             []guild.Webhook                      `json:"webhooks"`
}
