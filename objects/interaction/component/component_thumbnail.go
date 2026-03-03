package component

import (
	"encoding/json"
)

type Thumbnail struct {
	Id          *int              `json:"id,omitempty"`
	Media       UnfurledMediaItem `json:"media"`
	Description *string           `json:"description,omitempty"`
	Spoiler     *bool             `json:"spoiler,omitempty"`
}

type UnfurledMediaItem struct {
	Url         string  `json:"url"`
	ProxyUrl    *string `json:"proxy_url,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Width       *int    `json:"width,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
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
