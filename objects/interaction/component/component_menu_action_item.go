package component

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

// MenuActionItem is only valid inside an OverflowMenuGroup. Link style items
// use Url instead of CustomId.
type MenuActionItem struct {
	Id          *int                `json:"id,omitempty"`
	CustomId    string              `json:"custom_id,omitempty"`
	Label       string              `json:"label"`
	Style       MenuActionItemStyle `json:"style"`
	Description *string             `json:"description,omitempty"`
	Url         *string             `json:"url,omitempty"`
	Emoji       *emoji.Emoji        `json:"emoji,omitempty"`
	Disabled    bool                `json:"disabled,omitempty"`
}

type MenuActionItemStyle uint8

const (
	MenuActionItemStyleDefault MenuActionItemStyle = iota + 1
	MenuActionItemStyleDanger
	MenuActionItemStyleLink
)

func (m MenuActionItem) Type() ComponentType {
	return ComponentMenuActionItem
}

func (m MenuActionItem) MarshalJSON() ([]byte, error) {
	type WrappedMenuActionItem MenuActionItem

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedMenuActionItem
	}{
		Type:                  ComponentMenuActionItem,
		WrappedMenuActionItem: WrappedMenuActionItem(m),
	})
}

func BuildMenuActionItem(data MenuActionItem) Component {
	return Component{
		Type:          ComponentMenuActionItem,
		ComponentData: data,
	}
}
