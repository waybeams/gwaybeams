package display

type ComponentComposer func(b Builder)

type GlfwWindowHint int

const (
	AutoIconify = iota
	Decorated
	Floating
	Focused
	Iconified
	Maximized
	Resizable
	Visible
)

/*
// TODO(lbayes) Map local hints to glfw library hints
const (
	Focused     GlfwWindowHint = C.GLFW_FOCUSED      // Specifies whether the window will be given input focus when created. This hint is ignored for full screen and initially hidden windows.
	Iconified   Hint = C.GLFW_ICONIFIED    // Specifies whether the window will be minimized.
	Maximized   Hint = C.GLFW_MAXIMIZED    // Specifies whether the window is maximized.
	Visible     Hint = C.GLFW_VISIBLE      // Specifies whether the window will be initially visible.
	Resizable   Hint = C.GLFW_RESIZABLE    // Specifies whether the window will be resizable by the user.
	Decorated   Hint = C.GLFW_DECORATED    // Specifies whether the window will have window decorations such as a border, a close widget, etc.
	Floating    Hint = C.GLFW_FLOATING     // Specifies whether the window will be always-on-top.
	AutoIconify Hint = C.GLFW_AUTO_ICONIFY // Specifies whether fullscreen windows automatically iconify (and restore the previous video mode) on focus loss.
)
*/

const DefaultFrameRate = 60
const DefaultWindowWidth = 1024
const DefaultWindowHeight = 768
const DefaultWindowTitle = "Default Title"

type BuilderOption func(b Builder) error

func FrameRate(fps int) BuilderOption {
	return func(b Builder) error {
		b.FrameRate(fps)
		return nil
	}
}

func WindowSize(width int, height int) BuilderOption {
	return func(b Builder) error {
		b.WindowSize(width, height)
		return nil
	}
}

// WindowHints are how we configure GLFW windows
type windowHint struct {
	name  GlfwWindowHint
	value interface{}
}

func WindowHint(hintName GlfwWindowHint, value interface{}) BuilderOption {
	return func(b Builder) error {
		b.PushWindowHint(hintName, value)
		return nil
	}
}

func WindowTitle(title string) BuilderOption {
	return func(b Builder) error {
		b.WindowTitle(title)
		return nil
	}
}
