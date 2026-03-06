package message

type MessageApplication struct {
	Id          uint64  `json:"id,string"`
	CoverImage  *string `json:"cover_image,omitempty"`
	Description string  `json:"description"`
	Icon        *string `json:"icon,omitempty"`
	Name        string  `json:"name"`
}
