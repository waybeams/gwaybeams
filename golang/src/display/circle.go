package display

type Circle struct {
	Sprite
}

func NewCircle() Displayable {
	return &Circle{}
}
