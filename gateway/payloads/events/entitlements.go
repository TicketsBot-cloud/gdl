package events

import "github.com/TicketsBot-cloud/gdl/objects/entitlement"

type EntitlementCreate struct {
	entitlement.Entitlement
}

type EntitlementUpdate struct {
	entitlement.Entitlement
}

type EntitlementDelete struct {
	entitlement.Entitlement
}
