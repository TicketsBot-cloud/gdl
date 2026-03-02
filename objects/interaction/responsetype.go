package interaction

type ResponseType uint8

const (
	ResponseTypePong                                 ResponseType = 1
	ResponseTypeChannelMessageWithSource             ResponseType = 4
	ResponseTypeDeferredChannelMessageWithSource     ResponseType = 5
	ResponseTypeDeferredUpdateMessage                ResponseType = 6
	ResponseTypeUpdateMessage                        ResponseType = 7
	ResponseTypeApplicationCommandAutoCompleteResult ResponseType = 8
	ResponseTypeModal                                ResponseType = 9
	ResponseTypePremiumRequired                      ResponseType = 10 // Deprecated
	ResponseTypeLaunchActivity                       ResponseType = 12
)
