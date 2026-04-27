package message

import "github.com/TicketsBot-cloud/gdl/objects/guild/emoji"

type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        emoji.Emoji          `json:"emoji"`
	BurstColors  []string             `json:"burst_colors"`
}
