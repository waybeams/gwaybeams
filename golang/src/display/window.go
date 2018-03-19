package display

type window struct {
	Sprite
}

func NewWindow(opts *Opts, args ...interface{}) *window {
	decl, err := ProcessArgs(args)

	if err != nil {
		panic(err)
	}

	// Instantiate and configure the component
	win := &window{}
	win.Declaration(decl)
	return win
}
