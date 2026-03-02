package guild

import (
	"fmt"
	"strings"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
	"github.com/TicketsBot-cloud/gdl/objects/guild/sticker"
	"github.com/TicketsBot-cloud/gdl/objects/member"
)

type Guild struct {
	Id                          uint64                    `json:"id,string"`
	Name                        string                    `json:"name"`
	Icon                        *string                   `json:"icon"`
	IconHash                    *string                   `json:"icon_hash,omitempty"`
	Splash                      *string                   `json:"splash"`
	DiscoverySplash             *string                   `json:"discovery_splash"`
	Owner                       bool                      `json:"owner,omitempty"`
	OwnerId                     uint64                    `json:"owner_id,string"`
	Permissions                 uint64                    `json:"permissions,string,omitempty"`
	Region                      *string                   `json:"region,omitempty"`
	AfkChannelId                objects.NullableSnowflake `json:"afk_channel_id"`
	AfkTimeout                  int                       `json:"afk_timeout"`
	WidgetEnabled               *bool                     `json:"widget_enabled,omitempty"`
	WidgetChannelId             objects.NullableSnowflake `json:"widget_channel_id,omitempty"`
	VerificationLevel           int                       `json:"verification_level"`
	DefaultMessageNotifications int                       `json:"default_message_notifications"`
	ExplicitContentFilter       int                       `json:"explicit_content_filter"`
	Roles                       []Role                    `json:"roles"`
	Emojis                      []emoji.Emoji             `json:"emojis"`
	Features                    []GuildFeature            `json:"features"`
	MfaLevel                    int                       `json:"mfa_level"`
	ApplicationId               objects.NullableSnowflake `json:"application_id"`
	SystemChannelId             objects.NullableSnowflake `json:"system_channel_id"`
	SystemChannelFlags          uint16                    `json:"system_channel_flags"`
	RulesChannelId              objects.NullableSnowflake `json:"rules_channel_id,omitempty"`
	JoinedAt                    time.Time                 `json:"joined_at,omitempty"`
	Large                       bool                      `json:"large,omitempty"`
	Unavailable                 *bool                     `json:"unavailable,omitempty"`
	MemberCount                 int                       `json:"member_count,omitempty"`
	VoiceStates                 []VoiceState              `json:"voice_states,omitempty"`
	Members                     []member.Member           `json:"members,omitempty"`
	Channels                    []channel.Channel         `json:"channels,omitempty"`
	Threads                     []channel.Channel         `json:"threads,omitempty"`
	MaxPresences                *int                      `json:"max_presences,omitempty"`
	MaxMembers                  *int                      `json:"max_members,omitempty"`
	VanityUrlCode               *string                   `json:"vanity_url_code"`
	Description                 *string                   `json:"description"`
	Banner                      *string                   `json:"banner"`
	PremiumTier                 PremiumTier               `json:"premium_tier"`
	PremiumSubscriptionCount    *int                      `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string                    `json:"preferred_locale"`
	PublicUpdatesChannelId      objects.NullableSnowflake `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        *int                      `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   *int                      `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      int                       `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int                       `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen            `json:"welcome_screen,omitempty"`
	NsfwLevel                   int                       `json:"nsfw_level"`
	Stickers                    []sticker.Sticker         `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool                      `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelId       objects.NullableSnowflake `json:"safety_alerts_channel_id"`
	IncidentsData               *IncidentsData            `json:"incidents_data,omitempty"`
}

type IncidentsData struct {
	InvitesDisabledUntil *time.Time `json:"invites_disabled_until,omitempty"`
	DmsDisabledUntil     *time.Time `json:"dms_disabled_until,omitempty"`
	DmSpamDetectedAt     *time.Time `json:"dm_spam_detected_at,omitempty"`
	RaidDetectedAt       *time.Time `json:"raid_detected_at,omitempty"`
}

func (g *Guild) IconUrl() string {
	if g.Icon == nil || *g.Icon == "" {
		return ""
	}

	extension := "png"
	if strings.HasPrefix(*g.Icon, "a_") {
		extension = "gif"
	}

	return fmt.Sprintf("https://cdn.discordapp.com/icons/%d/%s.%s", g.Id, *g.Icon, extension)
}

func (g *Guild) ToCachedGuild() (cached CachedGuild) {
	cached = CachedGuild{
		Id:                          g.Id,
		Name:                        g.Name,
		Icon:                        g.Icon,
		Splash:                      g.Splash,
		Owner:                       g.Owner,
		OwnerId:                     g.OwnerId,
		Permissions:                 g.Permissions,
		Region:                      g.Region,
		AfkChannelId:                g.AfkChannelId,
		AfkTimeout:                  g.AfkTimeout,
		VerificationLevel:           g.VerificationLevel,
		DefaultMessageNotifications: g.DefaultMessageNotifications,
		ExplicitContentFilter:       g.ExplicitContentFilter,
		Features:                    g.Features,
		MfaLevel:                    g.MfaLevel,
		ApplicationId:               g.ApplicationId,
		WidgetEnabled:               g.WidgetEnabled,
		WidgetChannelId:             g.WidgetChannelId,
		SystemChannelId:             g.SystemChannelId,
		SystemChannelFlags:          g.SystemChannelFlags,
		RulesChannelId:              g.RulesChannelId,
		JoinedAt:                    g.JoinedAt,
		Large:                       g.Large,
		Unavailable:                 g.Unavailable,
		MemberCount:                 g.MemberCount,
		MaxPresences:                g.MaxPresences,
		MaxMembers:                  g.MaxMembers,
		VanityUrlCode:               g.VanityUrlCode,
		Description:                 g.Description,
		Banner:                      g.Banner,
		PremiumTier:                 g.PremiumTier,
		PremiumSubscriptionCount:    g.PremiumSubscriptionCount,
		PreferredLocale:             g.PreferredLocale,
		PublicUpdatesChannelId:      g.PublicUpdatesChannelId,
		MaxVideoChannelUsers:        g.MaxVideoChannelUsers,
	}

	for _, role := range g.Roles {
		cached.Roles = append(cached.Roles, role.Id)
	}

	for _, emoji := range g.Emojis {
		cached.Emojis = append(cached.Emojis, emoji.Id.Value)
	}

	for _, channel := range g.Channels {
		cached.Channels = append(cached.Channels, channel.Id)
	}

	return
}
