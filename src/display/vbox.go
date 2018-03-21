package display

type vbox struct {
	box
}

func (v *vbox) RenderChildren(s Surface) {
	// Throwaway brute force, primitive pass of fake flex layout
	flexWidthSum := 0.0
	for _, child := range v.children {
		flexWidthSum += child.GetFlexWidth()
	}
	flexWidthValue := v.GetWidth()
	for _, child := range v.children {
		child.Width(float64(child.GetFlexWidth()) * flexWidthValue)
	}

	flexHeightSum := 0.0
	for _, child := range v.children {
		flexHeightSum += child.GetFlexHeight()
	}
	flexHeightValue := v.GetHeight() / float64(flexHeightSum)

	var lastChild Displayable
	for _, child := range v.children {
		child.Height(float64(child.GetFlexHeight()) * flexHeightValue)
		if lastChild != nil {
			child.Y(lastChild.GetY() + lastChild.GetHeight())
		}
		lastChild = child
	}

	// Traverse the tree rendering children all the way down
	for _, child := range v.children {
		child.RenderChildren(s)
		child.Render(s)
	}
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
