package display

type window struct {
	Sprite
}

func Window(f Factory, args ...interface{}) {
	decl, err := ProcessArgs(args)

	if err != nil {
		panic(err)
	}

	// Instantiate and configure the component
	win := &window{}
	win.Declaration(decl)
	f.Push(win)
}
