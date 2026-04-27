package auditlog

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditEntryInfo struct {
	ApplicationId                 *uint64    `json:"application_id,string,omitempty"`             // APPLICATION_COMMAND_PERMISSION_UPDATE
	AutoModerationRuleName        string     `json:"auto_moderation_rule_name,omitempty"`         // AUTO_MODERATION_*
	AutoModerationRuleTriggerType string     `json:"auto_moderation_rule_trigger_type,omitempty"` // AUTO_MODERATION_*
	ChannelId                     *uint64    `json:"channel_id,string,omitempty"`                 // MEMBER_MOVE & MESSAGE_PIN & MESSAGE_UNPIN & MESSAGE_DELETE & STAGE_INSTANCE_*
	Count                         string     `json:"count,omitempty"`                             // MESSAGE_DELETE & MESSAGE_BULK_DELETE & MEMBER_DISCONNECT & MEMBER_MOVE
	DeleteMemberDays              string     `json:"delete_member_days,omitempty"`                // MEMBER_PRUNE
	Id                            *uint64    `json:"id,string,omitempty"`                         // MESSAGE_DELETE & MESSAGE_BULK_DELETE & MEMBER_DISCONNECT & MEMBER_MOVE
	IntegrationType               string     `json:"integration_type,omitempty"`                  // MEMBER_KICK & MEMBER_ROLE_UPDATE
	MembersRemoved                string     `json:"members_removed,omitempty"`                   // MEMBER_PRUNE
	MessageId                     *uint64    `json:"message_id,string,omitempty"`                 // MESSAGE_PIN & MESSAGE_UNPIN
	RoleName                      string     `json:"role_name,omitempty"`                         // CHANNEL_OVERWRITE_CREATE & CHANNEL_OVERWRITE_UPDATE & CHANNEL_OVERWRITE_DELETE
	Type                          EntityType `json:"type,omitempty"`                              // CHANNEL_OVERWRITE_CREATE & CHANNEL_OVERWRITE_UPDATE & CHANNEL_OVERWRITE_DELETE
	Status                        string     `json:"status,omitempty"`                            // VOICE_CHANNEL_STATUS_UPDATE
}
