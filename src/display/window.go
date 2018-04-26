package display

// Window is an outermost component that manages the application event loop.
// Concrete Window implementations will connect the component Draw() calls with
// an appropriate native rendering surface.
// I would like to remove this interface at some point.
type Window interface {
	Displayable

	Init()
	PollEvents() []Event
}
