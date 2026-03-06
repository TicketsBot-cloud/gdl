package events

import "github.com/TicketsBot-cloud/gdl/objects/guild/emoji"

type VoiceChannelEffectSend struct {
	ChannelId     uint64       `json:"channel_id,string"`
	GuildId       uint64       `json:"guild_id,string"`
	UserId        uint64       `json:"user_id,string"`
	Emoji         *emoji.Emoji `json:"emoji"`
	AnimationType *int         `json:"animation_type"`
	AnimationId   uint         `json:"animation_id"`
	SoundId       *uint64      `json:"sound_id,string"`
	SoundVolume   *float64     `json:"sound_volume"`
}
