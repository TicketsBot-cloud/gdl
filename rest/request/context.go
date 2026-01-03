package request

import "context"

type contextKey string

const auditReasonKey contextKey = "audit_reason"

// WithAuditReason adds an audit log reason to the context.
// This reason will be automatically included in the X-Audit-Log-Reason header
// for any Discord API requests made with this context.
func WithAuditReason(ctx context.Context, reason string) context.Context {
	return context.WithValue(ctx, auditReasonKey, reason)
}

// getAuditReason retrieves the audit log reason from the context.
// Returns an empty string if no reason is set.
func getAuditReason(ctx context.Context) string {
	if reason, ok := ctx.Value(auditReasonKey).(string); ok {
		return reason
	}
	return ""
}
