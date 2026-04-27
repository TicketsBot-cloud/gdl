package guild

import "github.com/TicketsBot-cloud/gdl/objects/member"

type CachedVoiceState struct {
	ChannelId uint64 `json:"channel_id"`
	SessionId string `json:"session_id"`
	Deaf      bool   `json:"deaf"`
	Mute      bool   `json:"mute"`
	SelfDeaf  bool   `json:"self_deaf"`
	SelfMute  bool   `json:"self_mute"`
	Suppress  bool   `json:"suppress"`
}

func (s *CachedVoiceState) ToVoiceState(guildId uint64, m member.Member) VoiceState {
	channelId := s.ChannelId
	var userId uint64
	if m.User != nil {
		userId = m.User.Id
	}

	return VoiceState{
		GuildId:   &guildId,
		ChannelId: &channelId,
		UserId:    userId,
		Member:    &m,
		SessionId: s.SessionId,
		Deaf:      s.Deaf,
		Mute:      s.Mute,
		SelfDeaf:  s.SelfDeaf,
		SelfMute:  s.SelfMute,
		Suppress:  s.Suppress,
	}
}
