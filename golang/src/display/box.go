package display

type Box struct {
	Sprite
}

func NewBox() Displayable {
	return &Box{}
}
