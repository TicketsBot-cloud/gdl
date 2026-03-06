package message

import (
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

// MessageInteraction is deprecated in favour of InteractionMetadata
type MessageInteraction struct {
	Id     uint64         `json:"id,string"`
	Type   int            `json:"type"`
	Name   string         `json:"name"`
	User   user.User      `json:"user"`
	Member *member.Member `json:"member,omitempty"`
}

type MessageInteractionMetadata struct {
	Id                            uint64                      `json:"id,string"`
	Type                          int                         `json:"type"`
	User                          user.User                   `json:"user"`
	AuthorizingIntegrationOwners  map[string]string           `json:"authorizing_integration_owners"`
	OriginalResponseMessageId     *uint64                     `json:"original_response_message_id,string,omitempty"`
	TargetUser                    *user.User                  `json:"target_user,omitempty"`
	TargetMessageId               *uint64                     `json:"target_message_id,string,omitempty"`
	InteractedMessageId           *uint64                     `json:"interacted_message_id,string,omitempty"`
	TriggeringInteractionMetadata *MessageInteractionMetadata `json:"triggering_interaction_metadata,omitempty"`
}
