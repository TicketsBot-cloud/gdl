package user

type Activity struct {
	Name          string       `json:"name"`
	Type          ActivityType `json:"type"`
	Url           *string      `json:"url,omitempty"`
	CreatedAt     int          `json:"created_at,omitempty"`
	Timestamps    *Timestamps  `json:"timestamps,omitempty"`
	ApplicationId uint64       `json:"application_id,string,omitempty"`
	Details       *string      `json:"details,omitempty"`
	State         *string      `json:"state,omitempty"`
	// TODO: Figure out how to handle emoji w/o import cycle
	Party    *Party           `json:"party,omitempty"`
	Assets   *Asset           `json:"assets,omitempty"`
	Secrets  *Secret          `json:"secrets,omitempty"`
	Instance *bool            `json:"instance,omitempty"`
	Flags    int              `json:"flags,omitempty"`
	Buttons  []ActivityButton `json:"buttons,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}
