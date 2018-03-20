package display

type application struct {
	vbox
	Surface Surface
}

func Application(args ...interface{}) *application {
	decl, err := NewDeclaration(args)
	if err != nil {
		panic(err)
	}

	instance := &application{}
	instance.Declaration(decl)
	return instance
}
