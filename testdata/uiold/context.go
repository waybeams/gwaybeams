package ui

import (
	"clock"
	"events"
	"font"
)

type Context interface {
	AddFont(name, path string)
	Builder() Builder
	Clock() clock.Clock
	Destroy()
	CreateFonts(s Surface)
	Font(name string) *font.Font
	Listen()
	OnFrameEntered(handler events.EventHandler) events.Unsubscriber
}
