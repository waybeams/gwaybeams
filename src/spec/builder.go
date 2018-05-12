package spec

import "clock"

type Builder interface {
	Clock() clock.Clock
	Listen()
	Root() ReadWriter
	Surface() Surface
	Window() Window
}
