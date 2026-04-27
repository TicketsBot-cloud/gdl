package events

type ChannelInfoChannel struct {
	Id             uint64  `json:"id,string"`
	Status         *string `json:"status,omitempty"`
	VoiceStartTime *int64  `json:"voice_start_time,omitempty"`
}

type ChannelInfo struct {
	GuildId  uint64               `json:"guild_id,string"`
	Channels []ChannelInfoChannel `json:"channels"`
}
