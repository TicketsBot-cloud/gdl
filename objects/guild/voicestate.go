package guild

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/member"
)

type VoiceState struct {
	GuildId                 *uint64        `json:"guild_id,string"`
	ChannelId               *uint64        `json:"channel_id,string"`
	UserId                  uint64         `json:"user_id,string"`
	Member                  *member.Member `json:"member,omitempty"`
	SessionId               string         `json:"session_id"`
	Deaf                    bool           `json:"deaf"`
	Mute                    bool           `json:"mute"`
	SelfDeaf                bool           `json:"self_deaf"`
	SelfMute                bool           `json:"self_mute"`
	SelfStream              bool           `json:"self_stream"`
	SelfVideo               bool           `json:"self_video"`
	Suppress                bool           `json:"suppress"`
	RequestToSpeakTimestamp *time.Time     `json:"request_to_speak_timestamp"`
}

func (s *VoiceState) ToCachedVoiceState() CachedVoiceState {
	var channelId uint64
	if s.ChannelId != nil {
		channelId = *s.ChannelId
	}

	return CachedVoiceState{
		ChannelId: channelId,
		SessionId: s.SessionId,
		Deaf:      s.Deaf,
		Mute:      s.Mute,
		SelfDeaf:  s.SelfDeaf,
		SelfMute:  s.SelfMute,
		Suppress:  s.Suppress,
	}
}
