package component

import (
	"encoding/json"
)

type ChannelSelect struct {
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    *bool          `json:"disabled"`
}

func (i ChannelSelect) Type() ComponentType {
	return ComponentChannelSelect
}

func (i ChannelSelect) MarshalJSON() ([]byte, error) {
	type WrappedChannelSelect ChannelSelect

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedChannelSelect
	}{
		Type:                 ComponentChannelSelect,
		WrappedChannelSelect: WrappedChannelSelect(i),
	})
}

func BuildChannelSelect(data ChannelSelect) Component {
	return Component{
		Type:          ComponentChannelSelect,
		ComponentData: data,
	}
}
