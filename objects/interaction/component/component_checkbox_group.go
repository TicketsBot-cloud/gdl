package component

import "encoding/json"

type CheckboxGroup struct {
	CustomId  string                 `json:"custom_id"`
	Options   []CheckboxGroupOption  `json:"options"`
	MinValues *int                   `json:"min_values,omitempty"`
	MaxValues *int                   `json:"max_values,omitempty"`
	Required  *bool                  `json:"required,omitempty"`
}

type CheckboxGroupOption struct {
	Value       string  `json:"value"`
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	Default     bool    `json:"default,omitempty"`
}

func (c CheckboxGroup) Type() ComponentType {
	return ComponentCheckboxGroup
}

func (c CheckboxGroup) MarshalJSON() ([]byte, error) {
	type WrappedCheckboxGroup CheckboxGroup

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedCheckboxGroup
	}{
		Type:                 ComponentCheckboxGroup,
		WrappedCheckboxGroup: WrappedCheckboxGroup(c),
	})
}

func BuildCheckboxGroup(data CheckboxGroup) Component {
	return Component{
		Type:          ComponentCheckboxGroup,
		ComponentData: data,
	}
}
