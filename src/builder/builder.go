package builder

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

type SurfaceType int

const (
	CairoSurface = iota
	ImageSurface
	FakeSurface
)

type Option func(b *builder) error

type windowHint struct {
	name  GlfwWindowHint
	value interface{}
}

type Builder interface {
	GetSurfaceType() SurfaceType
	GetFrameRate() int
	GetWindowHints() []*windowHint
	GetWindowHint(hintName GlfwWindowHint) interface{}
}

type builder struct {
	surfaceType SurfaceType
	frameRate   int
	windowHints []*windowHint
}

func (b *builder) applyDefaults() {
	b.frameRate = DefaultFrameRate
	b.applyDefaultWindowHints()
}

func (b *builder) applyDefaultWindowHints() {
	b.windowHints = []*windowHint{
		&windowHint{name: AutoIconify, value: false},
		&windowHint{name: Decorated, value: false},
		&windowHint{name: Floating, value: true},
		&windowHint{name: Focused, value: true},
		&windowHint{name: Iconified, value: false},
		&windowHint{name: Maximized, value: false},
		&windowHint{name: Resizable, value: true},
		&windowHint{name: Visible, value: true},
	}
}

func (b *builder) removeWindowHint(hintName GlfwWindowHint) {
	hints := b.windowHints
	for i := 0; i < len(hints); i++ {
		if hints[i].name == hintName {
			b.windowHints = append(hints[:i], hints[i+1:]...)
			return
		}
	}
}

func (b *builder) GetSurfaceType() SurfaceType {
	return b.surfaceType
}

func (b *builder) GetFrameRate() int {
	return b.frameRate
}

func (b *builder) GetWindowHint(hintName GlfwWindowHint) interface{} {
	for _, hint := range b.windowHints {
		if hint.name == hintName {
			return hint.value
		}
	}
	return nil
}

func (b *builder) GetWindowHints() []*windowHint {
	return b.windowHints
}

func NewBuilder(args ...Option) (Builder, error) {
	b := &builder{}
	b.applyDefaults()

	for _, arg := range args {
		err := arg(b)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

// Surface Option for Builder
func Surface(surfaceType SurfaceType) Option {
	return func(b *builder) error {
		b.surfaceType = surfaceType
		return nil
	}
}

func FrameRate(fps int) Option {
	return func(b *builder) error {
		b.frameRate = fps
		return nil
	}
}

func WindowHint(hintName GlfwWindowHint, value interface{}) Option {
	wHint := &windowHint{
		name:  hintName,
		value: value,
	}

	return func(b *builder) error {
		// First remove existing hint by same name if found
		b.removeWindowHint(hintName)

		b.windowHints = append(b.windowHints, wHint)
		return nil
	}
}
