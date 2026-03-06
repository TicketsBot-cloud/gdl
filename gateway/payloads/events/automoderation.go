package events

import "github.com/TicketsBot-cloud/gdl/objects/automoderation"

type AutoModerationRuleCreate struct {
	automoderation.Rule
}

type AutoModerationRuleUpdate struct {
	automoderation.Rule
}

type AutoModerationRuleDelete struct {
	automoderation.Rule
}

type AutoModerationActionExecution struct {
	GuildId              uint64                     `json:"guild_id,string"`
	Action               automoderation.Action      `json:"action"`
	RuleId               uint64                     `json:"rule_id,string"`
	RuleTriggerType      automoderation.TriggerType `json:"rule_trigger_type"`
	UserId               uint64                     `json:"user_id,string"`
	ChannelId            *uint64                    `json:"channel_id,string"`
	MessageId            *uint64                    `json:"message_id,string"`
	AlertSystemMessageId *uint64                    `json:"alert_system_message_id,string"`
	Content              *string                    `json:"content"`
	MatchedKeyword       *string                    `json:"matched_keyword"`
	MatchedContent       *string                    `json:"matched_content"`
}
