package component

import (
	"encoding/json"
)

type UserSelect struct {
	Id            *int                 `json:"id,omitempty"`
	CustomId      string               `json:"custom_id"`
	Placeholder   string               `json:"placeholder,omitempty"`
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	MinValues     *int                 `json:"min_values,omitempty"`
	MaxValues     *int                 `json:"max_values,omitempty"`
	Disabled      bool                 `json:"disabled,omitempty"`
	Required      *bool                `json:"required,omitempty"`
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
