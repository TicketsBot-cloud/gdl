package component

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

type SelectMenu struct {
	Id          *int           `json:"id,omitempty"`
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
	Required    *bool          `json:"required,omitempty"`
}

type SelectOption struct {
	Label       string       `json:"label"`
	Value       string       `json:"value"`
	Description *string      `json:"description,omitempty"`
	Emoji       *emoji.Emoji `json:"emoji,omitempty"`
	Default     bool         `json:"default,omitempty"`
}

type SelectDefaultValueType string

const (
	SelectDefaultValueTypeRole    SelectDefaultValueType = "role"
	SelectDefaultValueTypeUser    SelectDefaultValueType = "user"
	SelectDefaultValueTypeChannel SelectDefaultValueType = "channel"
)

type SelectDefaultValue struct {
	Id   uint64                 `json:"id,string"`
	Type SelectDefaultValueType `json:"type"`
}

func (s SelectMenu) Type() ComponentType {
	return ComponentStringSelect
}

func (s SelectMenu) MarshalJSON() ([]byte, error) {
	type WrappedSelectMenu SelectMenu

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedSelectMenu
	}{
		Type:              ComponentStringSelect,
		WrappedSelectMenu: WrappedSelectMenu(s),
	})
}

func BuildSelectMenu(data SelectMenu) Component {
	return Component{
		Type:          ComponentStringSelect,
		ComponentData: data,
	}
}
