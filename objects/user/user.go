package user

import (
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects"
)

type UserFlag int

const (
	UserFlagStaff                 UserFlag = 1 << 0
	UserFlagPartner               UserFlag = 1 << 1
	UserFlagHypesquad             UserFlag = 1 << 2
	UserFlagBugHunterLevel1       UserFlag = 1 << 3
	UserFlagHypesquadOnlineHouse1 UserFlag = 1 << 6
	UserFlagHypesquadOnlineHouse2 UserFlag = 1 << 7
	UserFlagHypesquadOnlineHouse3 UserFlag = 1 << 8
	UserFlagPremiumEarlySupporter UserFlag = 1 << 9
	UserFlagTeamPseudoUser        UserFlag = 1 << 10
	UserFlagBugHunterLevel2       UserFlag = 1 << 14
	UserFlagVerifiedBot           UserFlag = 1 << 16
	UserFlagVerifiedDeveloper     UserFlag = 1 << 17
	UserFlagCertifiedModerator    UserFlag = 1 << 18
	UserFlagBotHttpInteractions   UserFlag = 1 << 19
	UserFlagActiveDeveloper       UserFlag = 1 << 22
)

type PremiumType int

const (
	PremiumTypeNone         PremiumType = 0
	PremiumTypeNitroClassic PremiumType = 1
	PremiumTypeNitro        PremiumType = 2
	PremiumTypeNitroBasic   PremiumType = 3
)

type AvatarDecorationData struct {
	Asset string `json:"asset"`
	SkuId uint64 `json:"sku_id,string"`
}

type UserNameplate struct {
	SkuId   uint64 `json:"sku_id,string"`
	Asset   string `json:"asset"`
	Label   string `json:"label"`
	Palette string `json:"palette"`
}

type UserCollectibles struct {
	Nameplate *UserNameplate `json:"nameplate,omitempty"`
}

type UserPrimaryGuild struct {
	IdentityGuildId objects.NullableSnowflake `json:"identity_guild_id"`
	IdentityEnabled *bool                     `json:"identity_enabled,omitempty"`
	Tag             *string                   `json:"tag,omitempty"`
	Badge           *string                   `json:"badge,omitempty"`
}

type User struct {
	Id                   uint64                `json:"id,string"`
	Username             string                `json:"username"`
	Discriminator        Discriminator         `json:"discriminator"`
	GlobalName           *string               `json:"global_name"`
	Avatar               Avatar                `json:"avatar"`
	Bot                  bool                  `json:"bot,omitempty"`
	System               *bool                 `json:"system,omitempty"`
	MfaEnabled           *bool                 `json:"mfa_enabled,omitempty"`
	Banner               *string               `json:"banner,omitempty"`
	AccentColor          *int                  `json:"accent_color,omitempty"`
	Locale               *string               `json:"locale,omitempty"`
	Verified             *bool                 `json:"verified,omitempty"`
	Email                *string               `json:"email,omitempty"`
	Flags                *UserFlag             `json:"flags,omitempty"`
	PremiumType          *PremiumType          `json:"premium_type,omitempty"`
	PublicFlags          *UserFlag             `json:"public_flags,omitempty"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles         *UserCollectibles     `json:"collectibles,omitempty"`
	PrimaryGuild         *UserPrimaryGuild     `json:"primary_guild,omitempty"`
}

// shortcut, ignores errors
func (u *User) AvatarUrl(size int) string {
	hash := u.Avatar.String()
	// if blank avatar, return a blank string so that we can use omitempty
	if len(hash) == 0 {
		return ""
	}

	var extension string
	if u.Avatar.Animated {
		extension = "gif"
	} else {
		extension = "webp"
	}

	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.%s?size=%d", u.Id, u.Avatar.String(), extension, size)
}

func (u *User) EffectiveName() string {
	if u.GlobalName != nil && len(*u.GlobalName) > 0 {
		return *u.GlobalName
	}

	return u.Username
}

func (u *User) Mention() string {
	return fmt.Sprintf("<@%d>", u.Id)
}

func (u *User) ToCachedUser() CachedUser {
	return CachedUser{
		Username:   u.Username,
		GlobalName: u.GlobalName,
		Avatar:     u.Avatar.String(),
		Bot:        u.Bot,
	}
}
