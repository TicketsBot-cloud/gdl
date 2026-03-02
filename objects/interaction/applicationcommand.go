package interaction

type ApplicationCommand struct {
	Id                     uint64                     `json:"id,string,omitempty"`
	Type                   ApplicationCommandType     `json:"type"`
	ApplicationId          uint64                     `json:"application_id,string,omitempty"`
	GuildId                *uint64                    `json:"guild_id,string,omitempty"`
	Name                   string                     `json:"name"`
	Description            string                     `json:"description"`
	Options                []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                  `json:"default_member_permissions,omitempty"`
	Nsfw                   bool                       `json:"nsfw,omitempty"`
	Version                uint64                     `json:"version,string,omitempty"`
	Contexts               []InteractionContextType   `json:"contexts,omitempty"`
}

type ApplicationCommandType uint8

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
	ApplicationCommandTypePrimaryEntryPoint
)
