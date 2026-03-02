package integration

type IntegrationExpireBehaviour int

const (
	IntegrationExpireBehaviourRemoveRole IntegrationExpireBehaviour = 0
	IntegrationExpireBehaviourKick       IntegrationExpireBehaviour = 1
)
