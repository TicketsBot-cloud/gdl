package component

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

type Button struct {
	Label    string       `json:"label"`
	CustomId string       `json:"custom_id,omitempty"`
	Style    ButtonStyle  `json:"style"`
	Emoji    *emoji.Emoji `json:"emoji,omitempty"`
	SkuId    *uint64      `json:"sku_id,string,omitempty"`
	Url      *string      `json:"url,omitempty"`
	Disabled bool         `json:"disabled,omitempty"`

	// PanelId marks this button as a placeholder for a real panel's "Open Ticket" button
	// within a guild-authored Components V2 tree, rather than a literal button to send to
	// Discord. It is never sent to Discord as-is: a resolve pass always replaces a button
	// carrying PanelId with a fresh button built from the referenced panel's real
	// CustomId/Label/Emoji before the message goes out, clearing PanelId on the output.
	PanelId *int `json:"panel_id,omitempty"`
}

func (b Button) Type() ComponentType {
	return ComponentButton
}

type ButtonStyle uint8

const (
	ButtonStylePrimary ButtonStyle = iota + 1
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
	ButtonStylePremium
)

func (b Button) MarshalJSON() ([]byte, error) {
	type WrappedButton Button

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedButton
	}{
		Type:          ComponentButton,
		WrappedButton: WrappedButton(b),
	})
}

func BuildButton(data Button) Component {
	return Component{
		Type:          ComponentButton,
		ComponentData: data,
	}
}
