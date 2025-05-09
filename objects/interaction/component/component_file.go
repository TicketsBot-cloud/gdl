package component

import (
	"encoding/json"
)

type File struct {
	File    UnfurledMediaItem `json:"file"`
	Spoiler *bool             `json:"spoiler,omitempty"`
}

func (i File) Type() ComponentType {
	return ComponentFile
}

func (i File) MarshalJSON() ([]byte, error) {
	type WrappedFile File

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedFile
	}{
		Type:        ComponentFile,
		WrappedFile: WrappedFile(i),
	})
}

func BuildFile(data File) Component {
	return Component{
		Type:          ComponentFile,
		ComponentData: data,
	}
}
