package interaction

import "encoding/json"

type InteractionContextType uint8

const (
	InteractionContextGuild InteractionContextType = 0
	InteractionContextBotDM InteractionContextType = 1
	InteractionContextPrivateChannel InteractionContextType = 2
)

func (i InteractionContextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint8(i))
}

func (i *InteractionContextType) UnmarshalJSON(data []byte) error {
	var val uint8
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*i = InteractionContextType(val)
	return nil
}