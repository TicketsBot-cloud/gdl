package soundboard

import "github.com/TicketsBot-cloud/gdl/objects/user"

type SoundboardSound struct {
	SoundId   uint64     `json:"sound_id,string"`
	Name      string     `json:"name"`
	Volume    float64    `json:"volume"`
	EmojiId   *uint64    `json:"emoji_id,string"`
	EmojiName *string    `json:"emoji_name,omitempty"`
	GuildId   *uint64    `json:"guild_id,string,omitempty"`
	Available bool       `json:"available"`
	User      *user.User `json:"user,omitempty"`
}
