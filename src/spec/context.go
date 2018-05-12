package spec

import (
	"clock"
	"font"
)

type Context interface {
	Font(name string) *font.Font
	Clock() clock.Clock
}
