package embed

type EmbedImage struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}
