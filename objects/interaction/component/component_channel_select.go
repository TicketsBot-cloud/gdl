package component

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/channel"
)

type ChannelSelect struct {
	Id            *int                  `json:"id,omitempty"`
	CustomId      string                `json:"custom_id"`
	ChannelTypes  []channel.ChannelType `json:"channel_types,omitempty"`
	Placeholder   string                `json:"placeholder,omitempty"`
	DefaultValues []SelectDefaultValue  `json:"default_values,omitempty"`
	MinValues     *int                  `json:"min_values,omitempty"`
	MaxValues     *int                  `json:"max_values,omitempty"`
	Disabled      *bool                 `json:"disabled,omitempty"`
	Required      *bool                 `json:"required,omitempty"`
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
