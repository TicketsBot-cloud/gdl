package component

import (
	"encoding/json"
)

type MentionableSelect struct {
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   *int           `json:"min_values,omitempty"`
	MaxValues   *int           `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled"`
}

func (i MentionableSelect) Type() ComponentType {
	return ComponentMentionableSelect
}

func (i MentionableSelect) MarshalJSON() ([]byte, error) {
	type WrappedMentionableSelect MentionableSelect

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedMentionableSelect
	}{
		Type:                     ComponentMentionableSelect,
		WrappedMentionableSelect: WrappedMentionableSelect(i),
	})
}

func BuildMentionableSelect(data MentionableSelect) Component {
	return Component{
		Type:          ComponentMentionableSelect,
		ComponentData: data,
	}
}
