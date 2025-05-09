package events

import (
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type ThreadMembersUpdate struct {
	ThreadId         uint64                  `json:"id,string"`
	GuildId          uint64                  `json:"guild_id,string"`
	MemberCount      uint64                  `json:"member_count"`
	AddedMembers     []channel.ThreadMember  `json:"added_members"`
	RemovedMemberIds utils.Uint64StringSlice `json:"removed_member_ids"`
}
