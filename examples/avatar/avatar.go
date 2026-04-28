package main

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/cache"
	"github.com/TicketsBot-cloud/gdl/command"
	"github.com/TicketsBot-cloud/gdl/gateway"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/sirupsen/logrus"
)

func main() {
	cacheFactory := cache.MemoryCacheFactory(cache.CacheOptions{
		Users: true,
	})

	shardOptions := gateway.ShardOptions{
		ShardCount: gateway.ShardCount{
			Total:   1,
			Lowest:  0,
			Highest: 1,
		},
		RateLimitStore:     ratelimit.NewMemoryStore(),
		CacheFactory:       cacheFactory,
		GuildSubscriptions: false,
	}

	token := ""
	sm := gateway.NewShardManager(token, shardOptions)

	registerCommands(sm)

	sm.Connect()
	sm.WaitForInterrupt()
}

func registerCommands(sm *gateway.ShardManager) {
	// register a command handler with prefix !
	ch := command.NewCommandHandler(sm, "!")

	// create a new avatar command
	cmd := command.NewCommand("avatar", []string{"a"}, onCommand)

	// register our command
	ch.RegisterCommand(cmd)
}

func onCommand(ctx command.CommandContext) {
	if len(ctx.Mentions) == 0 {
		_, _ = ctx.Shard.CreateMessage(context.Background(), ctx.ChannelId, "You need to mention a user")
		return
	}

	for _, mention := range ctx.Mentions {
		_, err := ctx.Shard.CreateMessage(context.Background(), ctx.ChannelId, fmt.Sprintf("%s's avatar is: %s", mention.Username, mention.AvatarUrl(2048)))
		if err != nil {
			logrus.Warn(err.Error())
		}
	}
}
