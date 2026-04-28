package interaction

type InteractionContextType int8

const (
	InteractionContextGuild InteractionContextType = iota
	InteractionContextBotDM
	InteractionContextPrivateChannel
)
