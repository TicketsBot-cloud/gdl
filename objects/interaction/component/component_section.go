package component

import (
	"encoding/json"
)

type Section struct {
	Components []Component `json:"components"`
	Accessory  Component   `json:"accessory,omitempty"`
}

func (i Section) Type() ComponentType {
	return ComponentSection
}

func (i Section) MarshalJSON() ([]byte, error) {
	type WrappedSection Section

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedSection
	}{
		Type:           ComponentSection,
		WrappedSection: WrappedSection(i),
	})
}

func BuildSection(data Section) Component {
	return Component{
		Type:          ComponentSection,
		ComponentData: data,
	}
}
