package events

type VoiceChannelStatusUpdate struct {
	Id      uint64  `json:"id,string"`
	GuildId uint64  `json:"guild_id,string"`
	Status  *string `json:"status"`
}
