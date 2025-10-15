package interaction

type InteractionContextType int8

const (
	InteractionContextGuild InteractionContextType = 0
	InteractionContextBotDM InteractionContextType = 1
	InteractionContextPrivateChannel InteractionContextType = 2
)
