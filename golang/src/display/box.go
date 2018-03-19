package display

type box struct {
	Sprite
}

func Box(s Surface, args ...interface{}) *box {
	instance := NewBox()
	decl, _ := NewDeclaration(args)
	instance.Declaration(decl)
	return instance
}

func NewBox() *box {
	return &box{}
}
