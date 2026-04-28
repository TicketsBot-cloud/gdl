package message

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/channel/embed"
	"github.com/TicketsBot-cloud/gdl/objects/guild/sticker"
	"github.com/TicketsBot-cloud/gdl/objects/interaction/component"
	"github.com/TicketsBot-cloud/gdl/objects/member"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/utils"
)

type BaseThemeType int

const (
	BaseThemeUnset    BaseThemeType = 0
	BaseThemeDark     BaseThemeType = 1
	BaseThemeLight    BaseThemeType = 2
	BaseThemeDarker   BaseThemeType = 3
	BaseThemeMidnight BaseThemeType = 4
)

type SharedClientTheme struct {
	Colors        []string       `json:"colors"`
	GradientAngle int            `json:"gradient_angle"`
	BaseMix       int            `json:"base_mix"`
	BaseTheme     *BaseThemeType `json:"base_theme,omitempty"`
}

type Message struct {
	Id                   uint64                      `json:"id,string"`
	ChannelId            uint64                      `json:"channel_id,string"`
	GuildId              *uint64                     `json:"guild_id,string,omitempty"`
	Author               user.User                   `json:"author"`
	Member               *member.Member              `json:"member,omitempty"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp,omitempty"`
	Tts                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []MessageMentionedUser      `json:"mentions,omitempty"` // The user objects in the mentions array will only have the partial member field present in MESSAGE_CREATE and MESSAGE_UPDATE events from text-based guild channels
	MentionRoles         utils.Uint64StringSlice     `json:"mention_roles"`
	MentionChannels      []ChannelMention            `json:"mention_channels,omitempty"` // Not all channel mentions in a message will appear in mention_channels. Only textual channels that are visible to everyone in a lurkable guild will ever be included.
	Attachments          []channel.Attachment        `json:"attachments,omitempty"`
	Embeds               []embed.Embed               `json:"embeds,omitempty"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                interface{}                 `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookId            *uint64                     `json:"webhook_id,string,omitempty"`
	Type                 MessageType                 `json:"type"`
	Activity             *MessageActivity            `json:"activity,omitempty"`
	Application          *MessageApplication         `json:"application,omitempty"`
	ApplicationId        *uint64                     `json:"application_id,string,omitempty"`
	Flags                MessageFlag                 `json:"flags,omitempty"`
	MessageReference     *MessageReference           `json:"message_reference,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	InteractionMetadata  *MessageInteractionMetadata `json:"interaction_metadata,omitempty"`
	Interaction          *MessageInteraction         `json:"interaction,omitempty"` // Deprecated: use InteractionMetadata
	Thread               *channel.Channel            `json:"thread,omitempty"`
	Components           []component.Component       `json:"components,omitempty"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []sticker.Sticker           `json:"stickers,omitempty"`
	Position             *int                        `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData       `json:"role_subscription_data,omitempty"`
	Poll                 *Poll                       `json:"poll,omitempty"`
	Call                 *MessageCall                `json:"call,omitempty"`
	Resolved             json.RawMessage             `json:"resolved,omitempty"`
	SharedClientTheme    *SharedClientTheme          `json:"shared_client_theme,omitempty"`
}

var channelMentionRegex = regexp.MustCompile(`<#(\d+)>`)

func (m *Message) ChannelMentions() []uint64 {
	mentions := make([]uint64, 0)

	for _, match := range channelMentionRegex.FindAllStringSubmatch(m.Content, -1) {
		if len(match) < 2 {
			continue
		}

		raw := match[1]
		if id, err := strconv.ParseUint(raw, 10, 64); err == nil {
			mentions = append(mentions, id)
		}
	}

	return mentions
}
