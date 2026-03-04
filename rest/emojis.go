package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
	"github.com/TicketsBot-cloud/gdl/utils"
)

func ListGuildEmojis(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) ([]emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/emojis", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteListGuildEmojis, guildId),
		RateLimiter: rateLimiter,
	}

	var emojis []emoji.Emoji
	err, _ := endpoint.Request(ctx, token, nil, &emojis)
	return emojis, err
}

func GetGuildEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, emojiId uint64) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/emojis/%d", guildId, emojiId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetGuildEmoji, guildId),
		RateLimiter: rateLimiter,
	}

	var emoji emoji.Emoji
	if err, _ := endpoint.Request(ctx, token, nil, &emoji); err != nil {
		return emoji, err
	}

	return emoji, nil
}

type CreateEmojiData struct {
	Name  string
	Image Image
	Roles []uint64 // roles for which this emoji will be whitelisted
}

func CreateGuildEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, data CreateEmojiData) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/emojis", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteCreateGuildEmoji, guildId),
		RateLimiter: rateLimiter,
	}

	var emoji emoji.Emoji
	imageData, err := data.Image.Encode()
	if err != nil {
		return emoji, err
	}

	body := map[string]interface{}{
		"name":  data.Name,
		"image": imageData,
		"roles": utils.Uint64StringSlice(data.Roles),
	}

	if err, _ := endpoint.Request(ctx, token, body, &emoji); err != nil {
		return emoji, err
	}

	return emoji, nil
}

// updating Image is not permitted
func ModifyGuildEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, emojiId uint64, data CreateEmojiData) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/emojis/%d", guildId, emojiId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyGuildEmoji, guildId),
		RateLimiter: rateLimiter,
	}

	body := map[string]interface{}{
		"name":  data.Name,
		"roles": utils.Uint64StringSlice(data.Roles),
	}

	var emoji emoji.Emoji
	if err, _ := endpoint.Request(ctx, token, body, &emoji); err != nil {
		return emoji, err
	}

	return emoji, nil
}

func DeleteGuildEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, emojiId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/emojis/%d", guildId, emojiId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteGuildEmoji, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}

type ListApplicationEmojisResponse struct {
	Items []emoji.Emoji `json:"items"`
}

func ListApplicationEmojis(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64) ([]emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/emojis", applicationId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteListApplicationEmojis, applicationId),
		RateLimiter: rateLimiter,
	}

	var res ListApplicationEmojisResponse
	err, _ := endpoint.Request(ctx, token, nil, &res)
	return res.Items, err
}

func GetApplicationEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, emojiId uint64) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/emojis/%d", applicationId, emojiId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteGetApplicationEmoji, applicationId),
		RateLimiter: rateLimiter,
	}

	var e emoji.Emoji
	err, _ := endpoint.Request(ctx, token, nil, &e)
	return e, err
}

type CreateApplicationEmojiData struct {
	Name  string `json:"name"`
	Image Image  `json:"-"`
}

func CreateApplicationEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId uint64, data CreateApplicationEmojiData) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/applications/%d/emojis", applicationId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteCreateApplicationEmoji, applicationId),
		RateLimiter: rateLimiter,
	}

	var e emoji.Emoji
	imageData, err := data.Image.Encode()
	if err != nil {
		return e, err
	}

	body := map[string]interface{}{
		"name":  data.Name,
		"image": imageData,
	}

	err, _ = endpoint.Request(ctx, token, body, &e)
	return e, err
}

func ModifyApplicationEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, emojiId uint64, name string) (emoji.Emoji, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/applications/%d/emojis/%d", applicationId, emojiId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteModifyApplicationEmoji, applicationId),
		RateLimiter: rateLimiter,
	}

	body := map[string]interface{}{
		"name": name,
	}

	var e emoji.Emoji
	err, _ := endpoint.Request(ctx, token, body, &e)
	return e, err
}

func DeleteApplicationEmoji(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, applicationId, emojiId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/applications/%d/emojis/%d", applicationId, emojiId),
		Route:       ratelimit.NewApplicationRoute(ratelimit.RouteDeleteApplicationEmoji, applicationId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
