package display

import "fmt"

type box struct {
	Sprite
}

func Box(s Surface, args ...interface{}) *box {
	fmt.Println("Render Box!")
	instance := NewBox()
	instance.Render(s)
	return instance
}

func NewBox() *box {
	return &box{}
}
