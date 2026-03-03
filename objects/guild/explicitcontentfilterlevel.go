package guild

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled            ExplicitContentFilterLevel = 0
	ExplicitContentFilterLevelMembersWithoutRoles ExplicitContentFilterLevel = 1
	ExplicitContentFilterLevelAllMembers          ExplicitContentFilterLevel = 2
)
