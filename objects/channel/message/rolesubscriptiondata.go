package message

type RoleSubscriptionData struct {
	RoleSubscriptionListingId uint64 `json:"role_subscription_listing_id,string"`
	TierName                  string `json:"tier_name"`
	TotalMonthsSubscribed     int    `json:"total_months_subscribed"`
	IsRenewal                 bool   `json:"is_renewal"`
}
