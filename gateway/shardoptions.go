package gateway

import (
	"github.com/TicketsBot-cloud/gdl/cache"
	"github.com/TicketsBot-cloud/gdl/gateway/intents"
	"github.com/TicketsBot-cloud/gdl/objects/user"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
)

type ShardOptions struct {
	ShardCount           ShardCount
	CacheFactory         cache.CacheFactory
	RateLimitStore       ratelimit.RateLimitStore
	GuildSubscriptions   bool
	Presence             user.UpdateStatus
	Hooks                Hooks
	Debug                bool
	Intents              []intents.Intent
	LargeShardingBuckets int // defaults to 1. don't touch unless discord tell you to
}

type ShardCount struct {
	Total   int
	Lowest  int // Inclusive
	Highest int // Exclusive
}
