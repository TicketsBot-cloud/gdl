package channel

type PartialChannel struct {
	Id   uint64      `json:"id,string"`
	Type ChannelType `json:"type"`
	Name *string     `json:"name,omitempty"`
}
