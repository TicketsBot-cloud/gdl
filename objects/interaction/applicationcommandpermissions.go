package interaction

type ApplicationCommandPermissionType uint8

const (
	ApplicationCommandPermissionTypeRole    ApplicationCommandPermissionType = iota + 1
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)

type ApplicationCommandPermissions struct {
	Id         uint64                           `json:"id,string"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

type GuildApplicationCommandPermissions struct {
	Id            uint64                          `json:"id,string"`
	ApplicationId uint64                          `json:"application_id,string"`
	GuildId       uint64                          `json:"guild_id,string"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}
