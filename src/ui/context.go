package ui

import (
	"clock"
	"events"
)

type Context interface {
	Builder() Builder
	Clock() clock.Clock
	Destroy()
	Listen()
	OnFrameEntered(handler events.EventHandler) events.Unsubscriber
}
