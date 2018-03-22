package main

import . "display"

func CreateBoxesApp() (Displayable, error) {
	return NewBuilder(Title("Hello World"), Size(640, 480)).Build(func(b Builder) {
		Sprite(b)
		// Box(s, FlexWidth(1), FlexHeight(1), MaxWidth(321), MaxHeight(2423))
		// Box(b, &Opts{FlexWidth: 1, FlexHeight: 1, MaxWidth: 640, MaxHeight: 480})
		// Box(b, &Opts{FlexWidth: 1, FlexHeight: 1, MaxWidth: 320, MaxHeight: 280})
		// })
	})
}

func main() {
	_, err := CreateBoxesApp()
	if err != nil {
		panic(err)
	}

}
