package message

import "github.com/TicketsBot-cloud/gdl/objects/guild/emoji"

type Reaction struct {
	Count int
	Me    bool
	Emoji emoji.Emoji
}
