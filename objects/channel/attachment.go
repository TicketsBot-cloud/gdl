package channel

type Attachment struct {
	Id           uint64   `json:"id,string"`
	Filename     string   `json:"filename"`
	Title        *string  `json:"title,omitempty"`
	Description  *string  `json:"description,omitempty"`
	ContentType  *string  `json:"content_type,omitempty"`
	Size         int      `json:"size"`
	Url          string   `json:"url"`
	ProxyUrl     string   `json:"proxy_url"`
	Height       *int     `json:"height,omitempty"`
	Width        *int     `json:"width,omitempty"`
	Ephemeral    *bool    `json:"ephemeral,omitempty"`
	DurationSecs *float64 `json:"duration_secs,omitempty"`
	Waveform     *string  `json:"waveform,omitempty"`
	Flags        *int     `json:"flags,omitempty"`
}
