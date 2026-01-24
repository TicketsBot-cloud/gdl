package interaction

import (
	"encoding/json"

	"github.com/TicketsBot-cloud/gdl/objects/interaction/component"
)

type MessageComponentInteractionData struct {
	ComponentType component.ComponentType `json:"component_type"`
	IMessageComponentInteractionData
}

func (d MessageComponentInteractionData) AsButton() ButtonInteractionData {
	return d.IMessageComponentInteractionData.(ButtonInteractionData)
}

func (d MessageComponentInteractionData) AsSelectMenu() SelectMenuInteractionData {
	return d.IMessageComponentInteractionData.(SelectMenuInteractionData)
}

func (d MessageComponentInteractionData) AsFileUpload() FileUploadInteractionData {
	return d.IMessageComponentInteractionData.(FileUploadInteractionData)
}

func (d MessageComponentInteractionData) AsRadioGroup() RadioGroupInteractionData {
	return d.IMessageComponentInteractionData.(RadioGroupInteractionData)
}

func (d MessageComponentInteractionData) AsCheckboxGroup() CheckboxGroupInteractionData {
	return d.IMessageComponentInteractionData.(CheckboxGroupInteractionData)
}

func (d MessageComponentInteractionData) AsCheckbox() CheckboxInteractionData {
	return d.IMessageComponentInteractionData.(CheckboxInteractionData)
}

type IMessageComponentInteractionData interface {
	Type() component.ComponentType
}

type MessageComponentInteractionBaseData struct {
	ComponentType component.ComponentType `json:"component_type"`
	CustomId      string                  `json:"custom_id"`
}

type ButtonInteractionData struct {
	MessageComponentInteractionBaseData
}

func (d ButtonInteractionData) Type() component.ComponentType {
	return component.ComponentButton
}

type SelectMenuInteractionData struct {
	MessageComponentInteractionBaseData
	Values []string `json:"values"`
}

func (d SelectMenuInteractionData) Type() component.ComponentType {
	return component.ComponentSelectMenu
}

type FileUploadInteractionData struct {
	MessageComponentInteractionBaseData
	Values []uint64 `json:"values"`
}

func (d FileUploadInteractionData) Type() component.ComponentType {
	return component.ComponentFileUpload
}

type RadioGroupInteractionData struct {
	MessageComponentInteractionBaseData
	Value string `json:"value"`
}

func (d RadioGroupInteractionData) Type() component.ComponentType {
	return component.ComponentRadioGroup
}

type CheckboxGroupInteractionData struct {
	MessageComponentInteractionBaseData
	Values []string `json:"values"`
}

func (d CheckboxGroupInteractionData) Type() component.ComponentType {
	return component.ComponentCheckboxGroup
}

type CheckboxInteractionData struct {
	MessageComponentInteractionBaseData
	Value bool `json:"value"`
}

func (d CheckboxInteractionData) Type() component.ComponentType {
	return component.ComponentCheckbox
}

func (d *MessageComponentInteractionData) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var componentType component.ComponentType
	if rawType, ok := raw["component_type"].(float64); ok {
		componentType = component.ComponentType(rawType)
	} else {
		return component.ErrMissingType
	}

	var err error
	switch componentType {
	case component.ComponentActionRow:
		return component.ErrUnknownType
	case component.ComponentButton:
		var parsed ButtonInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	case component.ComponentSelectMenu:
		var parsed SelectMenuInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	case component.ComponentFileUpload:
		var parsed FileUploadInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	case component.ComponentRadioGroup:
		var parsed RadioGroupInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	case component.ComponentCheckboxGroup:
		var parsed CheckboxGroupInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	case component.ComponentCheckbox:
		var parsed CheckboxInteractionData
		err = json.Unmarshal(data, &parsed)
		d.IMessageComponentInteractionData = parsed
	default:
		return component.ErrUnknownType
	}

	if err != nil {
		return err
	}

	d.ComponentType = componentType
	return nil
}
