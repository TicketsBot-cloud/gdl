package component

import "encoding/json"

type RadioGroup struct {
	CustomId string             `json:"custom_id"`
	Options  []RadioGroupOption `json:"options"`
	Required *bool              `json:"required,omitempty"`
}

type RadioGroupOption struct {
	Value       string  `json:"value"`
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	Default     bool    `json:"default,omitempty"`
}

func (r RadioGroup) Type() ComponentType {
	return ComponentRadioGroup
}

func (r RadioGroup) MarshalJSON() ([]byte, error) {
	type WrappedRadioGroup RadioGroup

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedRadioGroup
	}{
		Type:              ComponentRadioGroup,
		WrappedRadioGroup: WrappedRadioGroup(r),
	})
}

func BuildRadioGroup(data RadioGroup) Component {
	return Component{
		Type:          ComponentRadioGroup,
		ComponentData: data,
	}
}
