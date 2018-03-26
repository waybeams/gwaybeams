package display

import "fmt"

type LabelComponent struct {
	Component
}

func (l *LabelComponent) Draw(surface Surface) {
	fmt.Println("Label.Draw called!")
}

var Label = NewComponentFactoryFrom(Box)
