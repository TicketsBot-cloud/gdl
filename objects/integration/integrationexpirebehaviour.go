package integration

type IntegrationExpireBehaviour int

const (
	IntegrationExpireBehaviourRemoveRole IntegrationExpireBehaviour = iota
	IntegrationExpireBehaviourKick
)
