package spec

import (
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/font"
)

type Context interface {
	Font(name string) *font.Font
	Clock() clock.Clock
}
