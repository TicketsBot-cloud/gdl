package component

import (
	"encoding/json"
)

type MediaGallery struct {
	Items []MediaGalleryItem `json:"items"`
}

type MediaGalleryItem struct {
	Media       UnfurledMediaItem `json:"media,omitempty"`
	Description *string           `json:"description,omitempty"`
	Spoiler     *bool             `json:"spoiler,omitempty"`
}

func (i MediaGallery) Type() ComponentType {
	return ComponentMediaGallery
}

func (i MediaGallery) MarshalJSON() ([]byte, error) {
	type WrappedMediaGallery MediaGallery

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedMediaGallery
	}{
		Type:                ComponentMediaGallery,
		WrappedMediaGallery: WrappedMediaGallery(i),
	})
}

func BuildMediaGallery(data MediaGallery) Component {
	return Component{
		Type:          ComponentMediaGallery,
		ComponentData: data,
	}
}
