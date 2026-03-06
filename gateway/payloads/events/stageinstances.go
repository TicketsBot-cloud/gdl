package events

import "github.com/TicketsBot-cloud/gdl/objects/guild/stage"

type StageInstanceCreate struct {
	stage.StageInstance
}

type StageInstanceUpdate struct {
	stage.StageInstance
}

type StageInstanceDelete struct {
	stage.StageInstance
}
