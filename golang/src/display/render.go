package display

func Render(surface Surface, node Displayable) error {
	node.Render(surface)
	return nil
}
