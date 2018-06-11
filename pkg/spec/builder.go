package spec

import "github.com/waybeams/waybeams/pkg/clock"

// Builder creates and configures the graphical environment.
type Builder interface {
	Clock() clock.Clock
	Listen()
	Root() ReadWriter
	Surface() Surface
	Window() Window
}
