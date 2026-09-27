package channel

type ChannelFlag uint64

const (
	ChannelFlagPinned                   ChannelFlag = 1 << 1
	ChannelFlagRequireTag               ChannelFlag = 1 << 4
	ChannelFlagHideMediaDownloadOptions ChannelFlag = 1 << 15
	ChannelFlagObfuscated               ChannelFlag = 1 << 17
	ChannelFlagIsSpoilerChannel         ChannelFlag = 1 << 21
)

func (f ChannelFlag) Has(flag ChannelFlag) bool {
	return f&flag != 0
}
