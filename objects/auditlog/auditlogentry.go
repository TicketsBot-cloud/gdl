package auditlog

type AuditLogEntry struct {
	TargetId   *string          `json:"target_id"`
	Changes    []AuditLogChange `json:"changes,omitempty"`
	UserId     *uint64          `json:"user_id,string,omitempty"`
	Id         uint64           `json:"id,string"`
	ActionType AuditLogEvent    `json:"action_type"`
	Options    *AuditEntryInfo  `json:"options,omitempty"`
	Reason     *string          `json:"reason,omitempty"`
}
