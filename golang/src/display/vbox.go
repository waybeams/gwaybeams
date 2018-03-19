package display

type vbox struct {
	box
}

func VBox(S Surface, args ...interface{}) *vbox {
	instance := NewVBox()
	decl, _ := NewDeclaration(args)
	instance.Declaration(decl)
	return instance
}

func NewVBox() *vbox {
	return &vbox{}
}
