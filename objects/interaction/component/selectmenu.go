package component

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

type SelectMenu struct {
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled"`
	Required    *bool          `json:"required,omitempty"`

	// IsPanelSelect marks this select menu as a placeholder for the multi-panel's
	// dropdown within a guild-authored Components V2 tree, rather than a literal select
	// menu to send to Discord. It is never sent to Discord as-is: a resolve pass always
	// replaces a select menu carrying IsPanelSelect with a fresh select menu built from
	// the multi-panel's unplaced panels before the message goes out, clearing
	// IsPanelSelect on the output.
	IsPanelSelect bool `json:"is_panel_select,omitempty"`
}

type SelectOption struct {
	Label       string       `json:"label"`
	Value       string       `json:"value"`
	Description *string      `json:"description,omitempty"`
	Emoji       *emoji.Emoji `json:"emoji,omitempty"`
	Default     bool         `json:"default"`
}

func (s SelectMenu) Type() ComponentType {
	return ComponentSelectMenu
}

func (s SelectMenu) MarshalJSON() ([]byte, error) {
	type WrappedSelectMenu SelectMenu

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedSelectMenu
	}{
		Type:              ComponentSelectMenu,
		WrappedSelectMenu: WrappedSelectMenu(s),
	})
}

func BuildSelectMenu(data SelectMenu) Component {
	return Component{
		Type:          ComponentSelectMenu,
		ComponentData: data,
	}
}
