package display

// ApplicationComponent belongs at the root of any component tree that will
// manage change over time.
type ApplicationComponent struct {
	Component
}

func NewApplication() Displayable {
	return &ApplicationComponent{}
}

var Application = NewComponentFactory("Application", NewApplication)
