package channel

type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
	ChannelTypeGuildNewsThread ChannelType = iota + 3
	ChannelTypeGuildPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
	ChannelTypeGuildDirectory
	ChannelTypeGuildForum
	ChannelTypeGuildMedia
)

func (t ChannelType) IsThread() bool {
	return t == ChannelTypeGuildPublicThread ||
		t == ChannelTypeGuildPrivateThread ||
		t == ChannelTypeGuildNewsThread
}

func (t ChannelType) IsPostBased() bool {
	return t == ChannelTypeGuildForum || t == ChannelTypeGuildMedia
}
