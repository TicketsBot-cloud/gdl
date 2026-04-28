package user

type ClientStatusType string

const (
	ClientStatusTypeOnline       ClientStatusType = "online"
	ClientStatusTypeDoNotDisturb ClientStatusType = "dnd"
	ClientStatusTypeIdle         ClientStatusType = "idle"
	ClientStatusTypeInvisible    ClientStatusType = "invisible"
	ClientStatusTypeOffline      ClientStatusType = "offline"
)
