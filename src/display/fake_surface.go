package display

import "errors"

type SurfaceCommand struct {
	Name string
	Args []interface{}
}

type FakeSurface struct {
	commands []SurfaceCommand
}

func (s *FakeSurface) GetCommands() []SurfaceCommand {
	return s.commands
}

func (s *FakeSurface) MoveTo(x float64, y float64) {
	args := []interface{}{x, y}
	s.commands = append(s.commands, SurfaceCommand{Name: "MoveTo", Args: args})
}

func (s *FakeSurface) SetRgba(r, g, b, a float64) {
	args := []interface{}{r, g, b, a}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetRgba", Args: args})
}

func (s *FakeSurface) SetLineWidth(width float64) {
	args := []interface{}{width}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetLineWidth", Args: args})
}

func (s *FakeSurface) Stroke() {
	s.commands = append(s.commands, SurfaceCommand{Name: "Stroke"})
}

func (s *FakeSurface) Arc(xc, yc, radius, angle1, angle2 float64) {
	args := []interface{}{xc, yc, radius, angle1, angle2}
	s.commands = append(s.commands, SurfaceCommand{Name: "Arc", Args: args})
}

func (s *FakeSurface) DrawRectangle(x, y, width, height float64) {
	args := []interface{}{x, y, width, height}
	s.commands = append(s.commands, SurfaceCommand{Name: "DrawRectangle", Args: args})
}

func (s *FakeSurface) Fill() {
	s.commands = append(s.commands, SurfaceCommand{Name: "Fill"})
}

func (s *FakeSurface) FillPreserve() {
	s.commands = append(s.commands, SurfaceCommand{Name: "FillPreserve"})
}

func (s *FakeSurface) Push(d Displayable) error {
	return errors.New("Unsupported method")
}

func (s *FakeSurface) GetRoot() Displayable {
	// Not sure how to throw when error is not part of the interface. :-(
	panic("Unsupported method")
}
