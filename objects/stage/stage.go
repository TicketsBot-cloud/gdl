package stage

type StageInstance struct {
	Id                    uint64  `json:"id,string"`
	GuildId               uint64  `json:"guild_id,string"`
	ChannelId             uint64  `json:"channel_id,string"`
	Topic                 string  `json:"topic"`
	PrivacyLevel          int     `json:"privacy_level"`
	DiscoverableDisabled  bool    `json:"discoverable_disabled"`
	GuildScheduledEventId *uint64 `json:"guild_scheduled_event_id,string"`
}
