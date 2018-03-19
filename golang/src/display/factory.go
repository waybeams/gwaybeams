package display

type Factory struct {
	stack Stack
}

func (f *Factory) getStack() Stack {
	if f.stack == nil {
		f.stack = NewStack()
	}
	return f.stack
}

func (f *Factory) Push(d Displayable) {
	f.getStack().Push(d)
}
