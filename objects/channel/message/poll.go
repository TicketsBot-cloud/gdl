package message

import (
	"time"

	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

type Poll struct {
	Question         PollMedia    `json:"question"`
	Answers          []PollAnswer `json:"answers"`
	Expiry           *time.Time   `json:"expiry"`
	AllowMultiselect bool         `json:"allow_multiselect"`
	LayoutType       int          `json:"layout_type"`
	Results          *PollResults `json:"results,omitempty"`
}

type PollMedia struct {
	Text  *string      `json:"text,omitempty"`
	Emoji *emoji.Emoji `json:"emoji,omitempty"`
}

type PollAnswer struct {
	AnswerId  int       `json:"answer_id,omitempty"`
	PollMedia PollMedia `json:"poll_media"`
}

type PollResults struct {
	IsFinalized  bool              `json:"is_finalized"`
	AnswerCounts []PollAnswerCount `json:"answer_counts"`
}

type PollAnswerCount struct {
	Id      int  `json:"id"`
	Count   int  `json:"count"`
	MeVoted bool `json:"me_voted"`
}
