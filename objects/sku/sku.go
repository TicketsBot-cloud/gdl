package sku

type SKUType int

const (
	SKUTypeDurable           SKUType = 2
	SKUTypeConsumable        SKUType = 3
	SKUTypeSubscription      SKUType = 5
	SKUTypeSubscriptionGroup SKUType = 6
)

type SKUFlag uint32

const (
	SKUFlagAvailable         SKUFlag = 1 << 2
	SKUFlagGuildSubscription SKUFlag = 1 << 7
	SKUFlagUserSubscription  SKUFlag = 1 << 8
)

type SKU struct {
	Id            uint64  `json:"id,string"`
	Type          SKUType `json:"type"`
	ApplicationId uint64  `json:"application_id,string"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Flags         SKUFlag `json:"flags"`
}
