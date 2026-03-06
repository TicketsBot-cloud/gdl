package application

type ActivityLocation struct {
	Id        string  `json:"id"`
	Kind      string  `json:"kind"`
	ChannelId uint64  `json:"channel_id,string"`
	GuildId   *uint64 `json:"guild_id,string,omitempty"`
}

type ActivityInstance struct {
	ApplicationId uint64           `json:"application_id,string"`
	InstanceId    string           `json:"instance_id"`
	LaunchId      string           `json:"launch_id"`
	Location      ActivityLocation `json:"location"`
	Users         []uint64         `json:"users"`
}
