package component

import (
	"encoding/json"
)

type TextDisplay struct {
	Content string `json:"content"`
}

func (i TextDisplay) Type() ComponentType {
	return ComponentTextDisplay
}

func (i TextDisplay) MarshalJSON() ([]byte, error) {
	type WrappedTextDisplay TextDisplay

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedTextDisplay
	}{
		Type:               ComponentTextDisplay,
		WrappedTextDisplay: WrappedTextDisplay(i),
	})
}

func BuildTextDisplay(data TextDisplay) Component {
	return Component{
		Type:          ComponentTextDisplay,
		ComponentData: data,
	}
}
