package component

import (
	"encoding/json"
)

type Label struct {
	Id          *int      `json:"id,omitempty"`
	Label       string    `json:"label"`
	Description *string   `json:"description,omitempty"`
	Component   Component `json:"component"`
}

func (i Label) Type() ComponentType {
	return ComponentLabel
}

func (i Label) MarshalJSON() ([]byte, error) {
	type WrappedLabel Label

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedLabel
	}{
		Type:         ComponentLabel,
		WrappedLabel: WrappedLabel(i),
	})
}

func BuildLabel(data Label) Component {
	return Component{
		Type:          ComponentLabel,
		ComponentData: data,
	}
}
