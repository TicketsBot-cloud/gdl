package component

import (
	"encoding/json"
)

type Separator struct {
	Divider *bool `json:"divider,omitempty"`
	Spacing *int
}

func (i Separator) Type() ComponentType {
	return ComponentSeparator
}

func (i Separator) MarshalJSON() ([]byte, error) {
	type WrappedSeparator Separator

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedSeparator
	}{
		Type:             ComponentSeparator,
		WrappedSeparator: WrappedSeparator(i),
	})
}

func BuildSeparator(data Separator) Component {
	return Component{
		Type:          ComponentSeparator,
		ComponentData: data,
	}
}
