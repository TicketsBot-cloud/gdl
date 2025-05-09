package component

import (
	"encoding/json"
)

type Thumbnail struct {
	Media       UnfurledMediaItem `json:"media,omitempty"`
	Description *string           `json:"description,omitempty"`
	Spoiler     *bool             `json:"spoiler,omitempty"`
}

type UnfurledMediaItem struct {
	Url         string `json:"url"`
	ProxyUrl    string `json:"proxy_url"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ContentType string `json:"content_type"`
}

func (i Thumbnail) Type() ComponentType {
	return ComponentThumbnail
}

func (i Thumbnail) MarshalJSON() ([]byte, error) {
	type WrappedThumbnail Thumbnail

	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		WrappedThumbnail
	}{
		Type:             ComponentThumbnail,
		WrappedThumbnail: WrappedThumbnail(i),
	})
}

func BuildThumbnail(data Thumbnail) Component {
	return Component{
		Type:          ComponentThumbnail,
		ComponentData: data,
	}
}
