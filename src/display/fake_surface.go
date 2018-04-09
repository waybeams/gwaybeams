package display

// SurfaceCommand stores method name and arguments for a given call.
type SurfaceCommand struct {
	Name string
	Args []interface{}
}

// FakeSurface is a drawing surface that is provided to test Draw implementations.
// Rather than rendering into some hardware interface, the methods provided here
// will simply record that they were called and with what arguments.
type FakeSurface struct {
	commands []SurfaceCommand
}

// GetCommands returns the collection of commands that have been made against
// a given instance of the FakeSurface.
func (s *FakeSurface) GetCommands() []SurfaceCommand {
	return s.commands
}

// SetFillColor stores the provided Hex RGBA fill color (e.g., 0xffcc00ff).
func (s *FakeSurface) SetFillColor(color uint) {
	args := []interface{}{color}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetFillColor", Args: args})
}

// SetStrokeColor stores the provided Hex RGBA stroke color (e.g., 0xffcc00ff).
func (s *FakeSurface) SetStrokeColor(color uint) {
	args := []interface{}{color}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetStrokeColor", Args: args})
}

func (s *FakeSurface) MoveTo(x float64, y float64) {
	args := []interface{}{x, y}
	s.commands = append(s.commands, SurfaceCommand{Name: "MoveTo", Args: args})
}

// SetStrokeWidth sets the stroke width
func (s *FakeSurface) SetStrokeWidth(width float64) {
	args := []interface{}{width}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetLineWidth", Args: args})
}

// Stroke draws a stroke around the last created shape.
func (s *FakeSurface) Stroke() {
	s.commands = append(s.commands, SurfaceCommand{Name: "Stroke"})
}

// Arc draws a arc along the provided point, radius and angles.
func (s *FakeSurface) Arc(xc, yc, radius, angle1, angle2 float64) {
	args := []interface{}{xc, yc, radius, angle1, angle2}
	s.commands = append(s.commands, SurfaceCommand{Name: "Arc", Args: args})
}

func (s *FakeSurface) BeginPath() {
	s.commands = append(s.commands, SurfaceCommand{Name: "BeginPath"})
}

func (s *FakeSurface) DebugDumpPathCache() {
	s.commands = append(s.commands, SurfaceCommand{Name: "DebugDumpCachePath"})
}

// Rect draws a rectangle on the provided point and width and height.
func (s *FakeSurface) Rect(x, y, width, height float64) {
	args := []interface{}{x, y, width, height}
	s.commands = append(s.commands, SurfaceCommand{Name: "Rect", Args: args})
}

// Fill will fill the last created shape.
func (s *FakeSurface) Fill() {
	s.commands = append(s.commands, SurfaceCommand{Name: "Fill"})
}

func (s *FakeSurface) SetFontSize(size float64) {
	args := []interface{}{size}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetFontSize", Args: args})
}

func (s *FakeSurface) SetFontFace(face string) {
	args := []interface{}{face}
	s.commands = append(s.commands, SurfaceCommand{Name: "SetFontFace", Args: args})
}

func (s *FakeSurface) Text(x float64, y float64, text string) {
	args := []interface{}{x, y, text}
	s.commands = append(s.commands, SurfaceCommand{Name: "Text", Args: args})
}

// GetOffsetSurfaceFor will return a OffsetSurface for the provided Displayable.
func (s *FakeSurface) GetOffsetSurfaceFor(d Displayable) Surface {
	// Do not return an offset surface, we want to store and verify the unmodified inputs.
	return s
}

func NewFakeSurface() *FakeSurface {
	return &FakeSurface{}
}
