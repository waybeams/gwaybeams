package main

import . "display"

func CreateBoxesApp(title string) (Displayable, error) {
	return NewBuilder(WindowTitle(title), WindowSize(640, 480)).Build(func(b Builder) {
		Sprite(b, Children(func() {
			Sprite(b, FlexWidth(1), FlexHeight(1), MaxWidth(640), MaxHeight(480))
			Sprite(b, FlexWidth(1), FlexHeight(1), MaxWidth(320), MaxHeight(240))
		}))
	})
}

func main() {
	_, err := CreateBoxesApp("Boxes Example")
	if err != nil {
		panic(err)
	}
}
