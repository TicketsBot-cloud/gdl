package guild

import "fmt"

type RoleColors struct {
	PrimaryColor   int  `json:"primary_color"`
	SecondaryColor *int `json:"secondary_color,omitempty"`
	TertiaryColor  *int `json:"tertiary_color,omitempty"`
}

type Role struct {
	Id           uint64     `json:"id,string"`
	Name         string     `json:"name"`
	Color        int        `json:"color"`
	Colors       RoleColors `json:"colors"`
	Hoist        bool       `json:"hoist"`
	Icon         *string    `json:"icon"`
	UnicodeEmoji *string    `json:"unicode_emoji"`
	Position     int        `json:"position"`
	Permissions  uint64     `json:"permissions,string"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
	Flags        int        `json:"flags"`
	Tags         RoleTags   `json:"tags"`
}

type RoleTags struct {
	BotId         *uint64 `json:"bot_id,string,omitempty"`
	IntegrationId *uint64 `json:"integration_id,string,omitempty"`
	// PremiumSubscriber, AvailableForPurchase, GuildConnections: key present (even as null) = tag applies
	PremiumSubscriber     *bool   `json:"premium_subscriber,omitempty"`
	SubscriptionListingId *uint64 `json:"subscription_listing_id,string,omitempty"`
	AvailableForPurchase  *bool   `json:"available_for_purchase,omitempty"`
	GuildConnections      *bool   `json:"guild_connections,omitempty"`
}

func (r *Role) Mention() string {
	return fmt.Sprintf("<@&%d>", r.Id)
}

func (r *Role) ToCachedRole(guildId uint64) CachedRole {
	return CachedRole{
		GuildId:     guildId,
		Name:        r.Name,
		Color:       r.Color,
		Hoist:       r.Hoist,
		Position:    r.Position,
		Permissions: r.Permissions,
		Managed:     r.Managed,
		Mentionable: r.Mentionable,
	}
}
