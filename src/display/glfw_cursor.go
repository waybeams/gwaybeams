package display

type Cursor interface {
}

type cursor struct {
	window Window
}

func (c *cursor) UpdatePosition(x int, y int) {
}

func NewCursor(win Window) Cursor {
	return &cursor{window: win}
}
