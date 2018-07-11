package spec

import "github.com/waybeams/waybeams/pkg/clock"

// Scheduler manages Specification lifecycle and environment interactions.
type Scheduler interface {
	Clock() clock.Clock
	Close()
	Listen()
	Root() ReadWriter
}
