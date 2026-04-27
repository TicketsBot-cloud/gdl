package channel

import (
	"encoding/json"
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type AttachmentFlag int

const (
	AttachmentFlagIsClip      AttachmentFlag = 1 << 0
	AttachmentFlagIsThumbnail AttachmentFlag = 1 << 1
	AttachmentFlagIsRemix     AttachmentFlag = 1 << 2
	AttachmentFlagIsSpoiler   AttachmentFlag = 1 << 3
	AttachmentFlagIsAnimated  AttachmentFlag = 1 << 5
)

type Attachment struct {
	Id                 uint64          `json:"id,string"`
	Filename           string          `json:"filename"`
	Title              *string         `json:"title,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ContentType        *string         `json:"content_type,omitempty"`
	Size               int             `json:"size"`
	Url                string          `json:"url"`
	ProxyUrl           string          `json:"proxy_url"`
	Height             *int            `json:"height,omitempty"`
	Width              *int            `json:"width,omitempty"`
	Placeholder        *string         `json:"placeholder,omitempty"`
	PlaceholderVersion *int            `json:"placeholder_version,omitempty"`
	Ephemeral          *bool           `json:"ephemeral,omitempty"`
	DurationSecs       *float64        `json:"duration_secs,omitempty"`
	Waveform           *string         `json:"waveform,omitempty"`
	Flags              *AttachmentFlag `json:"flags,omitempty"`
	ClipParticipants   []user.User     `json:"clip_participants,omitempty"`
	ClipCreatedAt      *time.Time      `json:"clip_created_at,omitempty"`
	Application        json.RawMessage `json:"application,omitempty"`
}
