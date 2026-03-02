package channel

type PermissionOverwrite struct {
	Id    uint64                  `json:"id,string"`
	Type  PermissionOverwriteType `json:"type"`
	Allow uint64                  `json:"allow,string"`
	Deny  uint64                  `json:"deny,string"`
}
