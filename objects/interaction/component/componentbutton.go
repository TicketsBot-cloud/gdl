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
	SkuId    *uint64      `json:"sku_id,omitempty"`
	Url      *string      `json:"url,omitempty"`
	Disabled bool         `json:"disabled"`
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
