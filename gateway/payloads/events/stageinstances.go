package events

import "github.com/TicketsBot-cloud/gdl/objects/stage"

type StageInstanceCreate struct {
	stage.StageInstance
}

type StageInstanceUpdate struct {
	stage.StageInstance
}

type StageInstanceDelete struct {
	stage.StageInstance
}
