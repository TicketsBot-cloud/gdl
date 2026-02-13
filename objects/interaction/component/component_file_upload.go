package component

import "encoding/json"

type FileUpload struct {
	CustomId  string `json:"custom_id"`
	MinValues *int   `json:"min_values,omitempty"`
	MaxValues *int   `json:"max_values,omitempty"`
	Required  *bool  `json:"required,omitempty"`
}

func (f FileUpload) Type() ComponentType {
	return ComponentFileUpload
}

func (f FileUpload) MarshalJSON() ([]byte, error) {
	type WrappedFileUpload FileUpload

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedFileUpload
	}{
		Type:              ComponentFileUpload,
		WrappedFileUpload: WrappedFileUpload(f),
	})
}

func BuildFileUpload(data FileUpload) Component {
	return Component{
		Type:          ComponentFileUpload,
		ComponentData: data,
	}
}
