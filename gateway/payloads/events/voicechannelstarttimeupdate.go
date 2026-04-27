package events

type VoiceChannelStartTimeUpdate struct {
	Id             uint64 `json:"id,string"`
	GuildId        uint64 `json:"guild_id,string"`
	VoiceStartTime *int64 `json:"voice_start_time"`
}
