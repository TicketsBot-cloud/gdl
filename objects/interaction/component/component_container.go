package component

import (
	"encoding/json"
)

type Container struct {
	Components  []Component `json:"components"`
	AccentColor *int        `json:"accent_color,omitempty"`
	Spoiler     *bool       `json:"spoiler,omitempty"`
}

func (i Container) Type() ComponentType {
	return ComponentContainer
}

func (i Container) MarshalJSON() ([]byte, error) {
	type WrappedContainer Container

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedContainer
	}{
		Type:             ComponentContainer,
		WrappedContainer: WrappedContainer(i),
	})
}

func BuildContainer(data Container) Component {
	return Component{
		Type:          ComponentContainer,
		ComponentData: data,
	}
}
