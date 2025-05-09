package component

import (
	"encoding/json"
)

type RoleSelect struct {
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled"`
}

func (i RoleSelect) Type() ComponentType {
	return ComponentRoleSelect
}

func (i RoleSelect) MarshalJSON() ([]byte, error) {
	type WrappedRoleSelect RoleSelect

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedRoleSelect
	}{
		Type:              ComponentRoleSelect,
		WrappedRoleSelect: WrappedRoleSelect(i),
	})
}

func BuildRoleSelect(data RoleSelect) Component {
	return Component{
		Type:          ComponentRoleSelect,
		ComponentData: data,
	}
}
