package events

import "github.com/TicketsBot-cloud/gdl/objects/guild/soundboard"

type GuildSoundboardSoundCreate struct {
	soundboard.SoundboardSound
}

type GuildSoundboardSoundUpdate struct {
	soundboard.SoundboardSound
}

type GuildSoundboardSoundDelete struct {
	SoundId uint64 `json:"sound_id,string"`
	GuildId uint64 `json:"guild_id,string"`
}

type GuildSoundboardSoundsUpdate struct {
	GuildId        uint64                         `json:"guild_id,string"`
	SoundboardSounds []soundboard.SoundboardSound `json:"soundboard_sounds"`
}

type SoundboardSounds struct {
	GuildId        uint64                         `json:"guild_id,string"`
	SoundboardSounds []soundboard.SoundboardSound `json:"soundboard_sounds"`
}
