package events

import "github.com/TicketsBot-cloud/gdl/objects/subscription"

type SubscriptionCreate struct {
	subscription.Subscription
}

type SubscriptionUpdate struct {
	subscription.Subscription
}

type SubscriptionDelete struct {
	subscription.Subscription
}
