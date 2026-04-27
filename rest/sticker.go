package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/guild/sticker"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type ModifyGuildStickerData struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags        *string `json:"tags,omitempty"`
}

type StickerPack struct {
	Id             uint64            `json:"id,string"`
	Stickers       []sticker.Sticker `json:"stickers"`
	Name           string            `json:"name"`
	SkuId          uint64            `json:"sku_id,string"`
	CoverStickerId *uint64           `json:"cover_sticker_id,string,omitempty"`
	Description    string            `json:"description"`
	BannerAssetId  *uint64           `json:"banner_asset_id,string,omitempty"`
}

func GetSticker(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, stickerId uint64) (sticker.Sticker, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/stickers/%d", stickerId),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetSticker, stickerId),
		RateLimiter: rateLimiter,
	}

	var s sticker.Sticker
	err, _ := endpoint.Request(ctx, token, nil, &s)
	return s, err
}

type stickerPacksResponse struct {
	StickerPacks []StickerPack `json:"sticker_packs"`
}

func GetStickerPacks(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter) ([]StickerPack, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    "/sticker-packs",
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetStickerPacks, 0),
		RateLimiter: rateLimiter,
	}

	var res stickerPacksResponse
	if err, _ := endpoint.Request(ctx, token, nil, &res); err != nil {
		return nil, err
	}

	return res.StickerPacks, nil
}

func GetStickerPack(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, packId uint64) (StickerPack, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/sticker-packs/%d", packId),
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetStickerPack, packId),
		RateLimiter: rateLimiter,
	}

	var pack StickerPack
	err, _ := endpoint.Request(ctx, token, nil, &pack)
	return pack, err
}

func ListGuildStickers(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) ([]sticker.Sticker, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/stickers", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteListGuildStickers, guildId),
		RateLimiter: rateLimiter,
	}

	var stickers []sticker.Sticker
	if err, _ := endpoint.Request(ctx, token, nil, &stickers); err != nil {
		return nil, err
	}

	return stickers, nil
}

func GetGuildSticker(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, stickerId uint64) (sticker.Sticker, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/stickers/%d", guildId, stickerId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetGuildSticker, guildId),
		RateLimiter: rateLimiter,
	}

	var s sticker.Sticker
	err, _ := endpoint.Request(ctx, token, nil, &s)
	return s, err
}

func ModifyGuildSticker(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, stickerId uint64, data ModifyGuildStickerData) (sticker.Sticker, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/stickers/%d", guildId, stickerId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyGuildSticker, guildId),
		RateLimiter: rateLimiter,
	}

	var s sticker.Sticker
	if err, _ := endpoint.Request(ctx, token, data, &s); err != nil {
		return sticker.Sticker{}, err
	}

	return s, nil
}

func DeleteGuildSticker(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, stickerId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/stickers/%d", guildId, stickerId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteGuildSticker, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
