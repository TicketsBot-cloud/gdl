package component

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ComponentType uint8

const (
	ComponentActionRow ComponentType = iota + 1
	ComponentButton
	ComponentSelectMenu
	ComponentInputText
	ComponentUserSelect
	ComponentRoleSelect
	ComponentMentionableSelect
	ComponentChannelSelect
	ComponentSection
	ComponentTextDisplay
	ComponentThumbnail
	ComponentMediaGallery
	ComponentFile
	ComponentSeparator
	ComponentContainer ComponentType = iota + 3 // 14
	ComponentLabel
	ComponentFileUpload
	ComponentRadioGroup ComponentType = iota + 4 // 21
	ComponentCheckboxGroup
	ComponentCheckbox
)

type Component struct {
	Type ComponentType `json:"type"`
	ComponentData
}

type ComponentData interface {
	Type() ComponentType
}

var (
	ErrMissingType  = errors.New("component was missing type field")
	ErrUnknownType  = errors.New("component had unknown type")
	ErrTypeMismatch = errors.New("data did not match component type")
)

func (c Component) MarshalJSON() ([]byte, error) {
	return encode(c.ComponentData)
}

func encode(c ComponentData) (json.RawMessage, error) {
	switch v := c.(type) {
	case ActionRow:
		subComponents := make([]json.RawMessage, len(v.Components))
		for i, sub := range v.Components {
			var err error
			subComponents[i], err = encode(sub.ComponentData)
			if err != nil {
				return nil, err
			}
		}

		data := map[string]interface{}{
			"type":       ComponentActionRow,
			"components": subComponents,
		}
		return json.Marshal(data)
	case Button:
		return json.Marshal(v)
	case SelectMenu:
		return json.Marshal(v)
	case InputText:
		return json.Marshal(v)
	case UserSelect:
		return json.Marshal(v)
	case RoleSelect:
		return json.Marshal(v)
	case MentionableSelect:
		return json.Marshal(v)
	case ChannelSelect:
		return json.Marshal(v)
	case Section:
		return json.Marshal(v)
	case TextDisplay:
		return json.Marshal(v)
	case Thumbnail:
		return json.Marshal(v)
	case MediaGallery:
		return json.Marshal(v)
	case File:
		return json.Marshal(v)
	case Separator:
		return json.Marshal(v)
	case Container:
		return json.Marshal(v)
	case Label:
		return json.Marshal(v)
	case FileUpload:
		return json.Marshal(v)
	case RadioGroup:
		return json.Marshal(v)
	case CheckboxGroup:
		return json.Marshal(v)
	case Checkbox:
		return json.Marshal(v)
	default:
		fmt.Println(v)
		return nil, ErrUnknownType
	}
}

func (c *Component) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var componentType ComponentType
	if rawType, ok := raw["type"].(float64); ok {
		componentType = ComponentType(rawType)
	} else {
		return ErrMissingType
	}

	var err error
	switch componentType {
	case ComponentActionRow:
		var parsed ActionRow
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentButton:
		var parsed Button
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentSelectMenu:
		var parsed SelectMenu
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentInputText:
		var parsed InputText
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentUserSelect:
		var parsed UserSelect
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentRoleSelect:
		var parsed RoleSelect
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentMentionableSelect:
		var parsed MentionableSelect
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentChannelSelect:
		var parsed ChannelSelect
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentSection:
		var parsed Section
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentTextDisplay:
		var parsed TextDisplay
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentThumbnail:
		var parsed Thumbnail
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentMediaGallery:
		var parsed MediaGallery
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentFile:
		var parsed File
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentSeparator:
		var parsed Separator
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentContainer:
		var parsed Container
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentLabel:
		var parsed Label
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentFileUpload:
		var parsed FileUpload
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentRadioGroup:
		var parsed RadioGroup
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentCheckboxGroup:
		var parsed CheckboxGroup
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	case ComponentCheckbox:
		var parsed Checkbox
		err = json.Unmarshal(data, &parsed)
		c.ComponentData = parsed
	default:
		return errors.Join(ErrUnknownType, fmt.Errorf("unknown component type %d", componentType))
	}

	if err != nil {
		return err
	}

	c.Type = componentType
	return nil
}
