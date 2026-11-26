package component

import "encoding/json"

// OverflowMenu can be placed in an action row or used as a section accessory.
// Clicking it opens a client-side menu of grouped MenuActionItems.
type OverflowMenu struct {
	Id     *int                `json:"id,omitempty"`
	Groups []OverflowMenuGroup `json:"groups"`
}

type OverflowMenuGroup struct {
	Label      *string          `json:"label,omitempty"`
	Components []MenuActionItem `json:"components"`
}

func (o OverflowMenu) Type() ComponentType {
	return ComponentOverflowMenu
}

func (o OverflowMenu) MarshalJSON() ([]byte, error) {
	type WrappedOverflowMenu OverflowMenu

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedOverflowMenu
	}{
		Type:                ComponentOverflowMenu,
		WrappedOverflowMenu: WrappedOverflowMenu(o),
	})
}

func BuildOverflowMenu(data OverflowMenu) Component {
	return Component{
		Type:          ComponentOverflowMenu,
		ComponentData: data,
	}
}
