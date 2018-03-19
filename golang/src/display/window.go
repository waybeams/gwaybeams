package display

type window struct {
	vbox
}

func Window(s Surface, args ...interface{}) *window {
	decl, err := NewDeclaration(args)

	if err != nil {
		panic(err)
	}

	// Instantiate and configure the component
	win := &window{}
	win.Declaration(decl)
	s.Push(win)

	return win
}
