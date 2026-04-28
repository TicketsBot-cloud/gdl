package user

type Asset struct {
	InviteCoverImage string `json:"invite_cover_image,omitempty"`
	LargeImage       string `json:"large_image,omitempty"`
	LargeText        string `json:"large_text,omitempty"`
	LargeUrl         string `json:"large_url,omitempty"`
	SmallImage       string `json:"small_image,omitempty"`
	SmallText        string `json:"small_text,omitempty"`
	SmallUrl         string `json:"small_url,omitempty"`
}
