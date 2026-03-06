package component

import "encoding/json"

type Checkbox struct {
	Id          *int    `json:"id,omitempty"`
	CustomId    string  `json:"custom_id"`
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	Default     bool    `json:"default,omitempty"`
}

func (c Checkbox) Type() ComponentType {
	return ComponentCheckbox
}

func (c Checkbox) MarshalJSON() ([]byte, error) {
	type WrappedCheckbox Checkbox

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedCheckbox
	}{
		Type:            ComponentCheckbox,
		WrappedCheckbox: WrappedCheckbox(c),
	})
}

func BuildCheckbox(data Checkbox) Component {
	return Component{
		Type:          ComponentCheckbox,
		ComponentData: data,
	}
}
