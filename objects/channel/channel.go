package channel

import (
	"fmt"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type Channel struct {
	Id                            uint64                    `json:"id,string"`
	Type                          ChannelType               `json:"type"`
	GuildId                       *uint64                   `json:"guild_id,string,omitempty"`
	Position                      *int                      `json:"position,omitempty"`
	PermissionOverwrites          []PermissionOverwrite     `json:"permission_overwrites,omitempty"`
	Name                          *string                   `json:"name,omitempty"`
	Topic                         *string                   `json:"topic,omitempty"`
	Nsfw                          bool                      `json:"nsfw,omitempty"`
	LastMessageId                 objects.NullableSnowflake `json:"last_message_id"`
	Bitrate                       int                       `json:"bitrate,omitempty"`
	UserLimit                     int                       `json:"user_limit,omitempty"`
	RateLimitPerUser              int                       `json:"rate_limit_per_user,omitempty"`
	Recipients                    []user.User               `json:"recipients,omitempty"`
	Icon                          *string                   `json:"icon,omitempty"`
	OwnerId                       *uint64                   `json:"owner_id,string,omitempty"`
	ApplicationId                 *uint64                   `json:"application_id,string,omitempty"`
	Managed                       *bool                     `json:"managed,omitempty"`
	ParentId                      objects.NullableSnowflake `json:"parent_id,omitempty"`
	LastPinTimestamp              *time.Time                `json:"last_pin_timestamp,omitempty"`
	RtcRegion                     *string                   `json:"rtc_region,omitempty"`
	VideoQualityMode              VideoQualityMode          `json:"video_quality_mode,omitempty"`
	MessageCount                  int                       `json:"message_count,omitempty"`
	MemberCount                   int                       `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata           `json:"thread_metadata,omitempty"`
	Member                        *ThreadMember             `json:"member,omitempty"`
	DefaultAutoArchiveDuration    *int                      `json:"default_auto_archive_duration,omitempty"`
	Permissions                   *string                   `json:"permissions,omitempty"`
	Flags                         *int                      `json:"flags,omitempty"`
	TotalMessageSent              *int                      `json:"total_message_sent,omitempty"`
	AvailableTags                 []ForumTag                `json:"available_tags,omitempty"`
	AppliedTags                   utils.Uint64StringSlice   `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction          `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int                      `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *int                      `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *int                      `json:"default_forum_layout,omitempty"`
}

func (c *Channel) Mention() string {
	return fmt.Sprintf("<#%d>", c.Id)
}

func (c *Channel) ToCachedChannel() CachedChannel {
	var guildId uint64
	if c.GuildId != nil {
		guildId = *c.GuildId
	}

	return CachedChannel{
		GuildId:              guildId,
		Type:                 c.Type,
		Position:             c.Position,
		PermissionOverwrites: c.PermissionOverwrites,
		Name:                 c.Name,
		Topic:                c.Topic,
		Nsfw:                 c.Nsfw,
		LastMessageId:        c.LastMessageId,
		Bitrate:              c.Bitrate,
		UserLimit:            c.UserLimit,
		RateLimitPerUser:     c.RateLimitPerUser,
		Icon:                 c.Icon,
		OwnerId:              c.OwnerId,
		ApplicationId:        c.ApplicationId,
		ParentId:             c.ParentId,
		LastPinTimestamp:     c.LastPinTimestamp,
	}
}

func (c *Channel) ToPartialChannel() PartialChannel {
	return PartialChannel{
		Id:   c.Id,
		Type: c.Type,
		Name: c.Name,
	}
}
