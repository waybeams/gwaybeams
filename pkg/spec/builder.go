package spec

import "github.com/waybeams/waybeams/pkg/clock"

type Builder interface {
	Clock() clock.Clock
	Listen()
	Root() ReadWriter
	Surface() Surface
	Window() Window
}
