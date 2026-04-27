package message

type MessageFlag uint

const (
	FlagCrossposted MessageFlag = 1 << iota
	FlagIsCrosspost
	FlagSuppressEmbeds
	FlagSourceMessageDeleted
	FlagUrgent
	FlagHasThread
	FlagEphemeral
	FlagLoading
	FlagFailedToMentionSomeRolesInThread
)

const (
	FlagSuppressNotifications MessageFlag = 1 << (iota + 12)
	FlagIsVoiceMessage
	FlagHasSnapshot
	FlagIsComponentsV2
)

func SumFlags(flags ...MessageFlag) (sum uint) {
	for _, flag := range flags {
		sum += uint(flag)
	}

	return
}
