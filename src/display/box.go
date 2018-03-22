package display

type box struct {
	SpriteComponent
}

func Box(s Surface, args ...interface{}) *box {
	decl, err := NewDeclaration(args)
	if err != nil {
		panic(err)
	}

	instance := &box{}
	instance.Declaration(decl)
	s.Push(instance)
	return instance
}

func NewBox() *box {
	return &box{}
}
