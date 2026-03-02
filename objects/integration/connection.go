package integration

type ConnectionVisibility int

const (
	ConnectionVisibilityNone     ConnectionVisibility = 0
	ConnectionVisibilityEveryone ConnectionVisibility = 1
)

type Connection struct {
	Id           string               `json:"id"`
	Name         string               `json:"name"`
	Type         string               `json:"type"` // twitch, youtube, spotify, etc.
	Revoked      *bool                `json:"revoked,omitempty"`
	Integrations []Integration        `json:"integrations,omitempty"`
	Verified     bool                 `json:"verified"`
	FriendSync   bool                 `json:"friend_sync"`
	ShowActivity bool                 `json:"show_activity"`
	TwoWayLink   bool                 `json:"two_way_link"`
	Visibility   ConnectionVisibility `json:"visibility"`
}
