package events

type MessagePollVoteAdd struct {
	UserId    uint64  `json:"user_id,string"`
	ChannelId uint64  `json:"channel_id,string"`
	MessageId uint64  `json:"message_id,string"`
	GuildId   *uint64 `json:"guild_id,string"`
	AnswerId  uint8   `json:"answer_id"`
}

type MessagePollVoteRemove struct {
	UserId    uint64  `json:"user_id,string"`
	ChannelId uint64  `json:"channel_id,string"`
	MessageId uint64  `json:"message_id,string"`
	GuildId   *uint64 `json:"guild_id,string"`
	AnswerId  uint8   `json:"answer_id"`
}
