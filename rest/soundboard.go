package rest

import (
	"context"
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/soundboard"
	"github.com/TicketsBot-cloud/gdl/rest/ratelimit"
	"github.com/TicketsBot-cloud/gdl/rest/request"
)

type SendSoundboardSoundData struct {
	SoundId       uint64  `json:"sound_id,string"`
	SourceGuildId *uint64 `json:"source_guild_id,string,omitempty"`
}

type CreateGuildSoundboardSoundData struct {
	Name      string   `json:"name"`
	Sound     string   `json:"sound"`
	Volume    *float64 `json:"volume,omitempty"`
	EmojiId   *uint64  `json:"emoji_id,string,omitempty"`
	EmojiName *string  `json:"emoji_name,omitempty"`
}

type ModifyGuildSoundboardSoundData struct {
	Name      *string  `json:"name,omitempty"`
	Volume    *float64 `json:"volume,omitempty"`
	EmojiId   *uint64  `json:"emoji_id,string,omitempty"`
	EmojiName *string  `json:"emoji_name,omitempty"`
}

func SendSoundboardSound(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, channelId uint64, data SendSoundboardSoundData) error {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/channels/%d/send-soundboard-sound", channelId),
		Route:       ratelimit.NewChannelRoute(ratelimit.RouteSendSoundboardSound, channelId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, data, nil)
	return err
}

func GetDefaultSoundboardSounds(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter) ([]soundboard.SoundboardSound, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    "/soundboard-default-sounds",
		Route:       ratelimit.NewOtherRoute(ratelimit.RouteGetDefaultSoundboardSounds, 0),
		RateLimiter: rateLimiter,
	}

	var sounds []soundboard.SoundboardSound
	if err, _ := endpoint.Request(ctx, token, nil, &sounds); err != nil {
		return nil, err
	}

	return sounds, nil
}

func ListGuildSoundboardSounds(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64) ([]soundboard.SoundboardSound, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/soundboard-sounds", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteListGuildSoundboardSounds, guildId),
		RateLimiter: rateLimiter,
	}

	var sounds []soundboard.SoundboardSound
	if err, _ := endpoint.Request(ctx, token, nil, &sounds); err != nil {
		return nil, err
	}

	return sounds, nil
}

func GetGuildSoundboardSound(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, soundId uint64) (soundboard.SoundboardSound, error) {
	endpoint := request.Endpoint{
		RequestType: request.GET,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/soundboard-sounds/%d", guildId, soundId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteGetGuildSoundboardSound, guildId),
		RateLimiter: rateLimiter,
	}

	var sound soundboard.SoundboardSound
	err, _ := endpoint.Request(ctx, token, nil, &sound)
	return sound, err
}

func CreateGuildSoundboardSound(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId uint64, data CreateGuildSoundboardSoundData) (soundboard.SoundboardSound, error) {
	endpoint := request.Endpoint{
		RequestType: request.POST,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/soundboard-sounds", guildId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteCreateGuildSoundboardSound, guildId),
		RateLimiter: rateLimiter,
	}

	var sound soundboard.SoundboardSound
	if err, _ := endpoint.Request(ctx, token, data, &sound); err != nil {
		return soundboard.SoundboardSound{}, err
	}

	return sound, nil
}

func ModifyGuildSoundboardSound(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, soundId uint64, data ModifyGuildSoundboardSoundData) (soundboard.SoundboardSound, error) {
	endpoint := request.Endpoint{
		RequestType: request.PATCH,
		ContentType: request.ApplicationJson,
		Endpoint:    fmt.Sprintf("/guilds/%d/soundboard-sounds/%d", guildId, soundId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteModifyGuildSoundboardSound, guildId),
		RateLimiter: rateLimiter,
	}

	var sound soundboard.SoundboardSound
	if err, _ := endpoint.Request(ctx, token, data, &sound); err != nil {
		return soundboard.SoundboardSound{}, err
	}

	return sound, nil
}

func DeleteGuildSoundboardSound(ctx context.Context, token string, rateLimiter *ratelimit.Ratelimiter, guildId, soundId uint64) error {
	endpoint := request.Endpoint{
		RequestType: request.DELETE,
		ContentType: request.Nil,
		Endpoint:    fmt.Sprintf("/guilds/%d/soundboard-sounds/%d", guildId, soundId),
		Route:       ratelimit.NewGuildRoute(ratelimit.RouteDeleteGuildSoundboardSound, guildId),
		RateLimiter: rateLimiter,
	}

	err, _ := endpoint.Request(ctx, token, nil, nil)
	return err
}
