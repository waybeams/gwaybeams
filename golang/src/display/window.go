package display

type window struct {
	Sprite
}

/*
func Window(r Renderer, s Signals, opts *Opts, styles *Styles) {

}
*/

func NewWindow(opts *Opts, args ...interface{}) *window {
	decl, err := NewDeclaration(args)

	if err != nil {
		panic(err)
	}

	// Instantiate and configure the component
	win := &window{}
	win.Declaration(decl)
	return win
}
