package subscription

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/utils"
)

type SubscriptionStatus int

const (
	SubscriptionStatusActive SubscriptionStatus = iota
	SubscriptionStatusEnding
	SubscriptionStatusInactive
)

type Subscription struct {
	Id                 uint64                  `json:"id,string"`
	UserId             uint64                  `json:"user_id,string"`
	SkuIds             utils.Uint64StringSlice `json:"sku_ids"`
	EntitlementIds     utils.Uint64StringSlice `json:"entitlement_ids"`
	RenewalSkuIds      utils.Uint64StringSlice `json:"renewal_sku_ids,omitempty"`
	CurrentPeriodStart time.Time               `json:"current_period_start"`
	CurrentPeriodEnd   time.Time               `json:"current_period_end"`
	Status             SubscriptionStatus      `json:"status"`
	CanceledAt         *time.Time              `json:"canceled_at,omitempty"`
	Country            *string                 `json:"country,omitempty"`
}
