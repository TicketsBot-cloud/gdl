package component

import (
	"encoding/json"
)

type UserSelect struct {
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled"`
}

func (i UserSelect) Type() ComponentType {
	return ComponentUserSelect
}

func (i UserSelect) MarshalJSON() ([]byte, error) {
	type WrappedUserSelect UserSelect

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedUserSelect
	}{
		Type:              ComponentUserSelect,
		WrappedUserSelect: WrappedUserSelect(i),
	})
}

func BuildUserSelect(data UserSelect) Component {
	return Component{
		Type:          ComponentUserSelect,
		ComponentData: data,
	}
}
