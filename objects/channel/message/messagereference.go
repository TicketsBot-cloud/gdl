package message

type MessageReference struct {
	Type            *int    `json:"type,omitempty"`
	MessageId       *uint64 `json:"message_id,string,omitempty"`
	ChannelId       *uint64 `json:"channel_id,string,omitempty"`
	GuildId         *uint64 `json:"guild_id,string,omitempty"`
	FailIfNotExists *bool   `json:"fail_if_not_exists,omitempty"`
}
