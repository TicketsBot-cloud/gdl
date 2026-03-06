package interaction

import "github.com/TicketsBot-cloud/gdl/objects/application"

type ApplicationCommand struct {
	Id                       uint64                                   `json:"id,string,omitempty"`
	Type                     ApplicationCommandType                   `json:"type"`
	ApplicationId            uint64                                   `json:"application_id,string,omitempty"`
	GuildId                  *uint64                                  `json:"guild_id,string,omitempty"`
	Name                     string                                   `json:"name"`
	NameLocalizations        map[string]string                        `json:"name_localizations,omitempty"`
	Description              string                                   `json:"description"`
	DescriptionLocalizations map[string]string                        `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption               `json:"options,omitempty"`
	DefaultMemberPermissions *string                                  `json:"default_member_permissions,omitempty"`
	Nsfw                     bool                                     `json:"nsfw,omitempty"`
	IntegrationTypes         []application.ApplicationIntegrationType `json:"integration_types,omitempty"`
	Contexts                 []InteractionContextType                 `json:"contexts,omitempty"`
	Version                  uint64                                   `json:"version,string,omitempty"`
	Handler                  *ApplicationCommandHandlerType           `json:"handler,omitempty"`
}

type ApplicationCommandType uint8

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
	ApplicationCommandTypePrimaryEntryPoint
)

type ApplicationCommandHandlerType uint8

const (
	ApplicationCommandHandlerTypeAppHandler ApplicationCommandHandlerType = iota + 1
	ApplicationCommandHandlerTypeDiscordLaunchActivity
)
