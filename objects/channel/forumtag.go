package channel

type ForumTag struct {
	Id        uint64  `json:"id,string"`
	Name      string  `json:"name"`
	Moderated bool    `json:"moderated"`
	EmojiId   *uint64 `json:"emoji_id,string,omitempty"`
	EmojiName *string `json:"emoji_name,omitempty"`
}

type DefaultReaction struct {
	EmojiId   *uint64 `json:"emoji_id,string,omitempty"`
	EmojiName *string `json:"emoji_name,omitempty"`
}
