package interaction

import "github.com/TicketsBot-cloud/gdl/objects/channel"

type ApplicationCommandInteractionDataOption struct {
	Type    ApplicationCommandOptionType              `json:"type"`
	Name    string                                    `json:"name"`
	Value   interface{}                               `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused bool                                      `json:"focused,omitempty"`
}

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType     `json:"type"`
	Name                     string                           `json:"name"`
	NameLocalizations        map[string]string                `json:"name_localizations,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations map[string]string                `json:"description_localizations,omitempty"`
	Required                 bool                             `json:"required,omitempty"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Autocomplete             bool                             `json:"autocomplete,omitempty"`
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []channel.ChannelType            `json:"channel_types,omitempty"`
	MinValue                 *float64                         `json:"min_value,omitempty"`
	MaxValue                 *float64                         `json:"max_value,omitempty"`
	MinLength                *int                             `json:"min_length,omitempty"`
	MaxLength                *int                             `json:"max_length,omitempty"`
}

type ApplicationCommandOptionType uint8

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value             interface{}       `json:"value"` // string, int or double
}
