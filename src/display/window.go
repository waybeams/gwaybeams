package display

type window struct {
	vbox
}

func Window(s Surface, args ...interface{}) *window {
	decl, err := NewDeclaration(args)
	if err != nil {
		panic(err)
	}

	instance := &window{}
	instance.Declaration(decl)
	s.Push(instance)
	return instance
}
