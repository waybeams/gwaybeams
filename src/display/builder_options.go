package display

type GlfwBuilderOption func(b GlfwBuilder) error

func FrameRate(fps int) GlfwBuilderOption {
	return func(b GlfwBuilder) error {
		b.FrameRate(fps)
		return nil
	}
}

func WindowSize(width int, height int) GlfwBuilderOption {
	return func(b GlfwBuilder) error {
		b.WindowSize(width, height)
		return nil
	}
}

// WindowHints are how we configure GLFW windows
type windowHint struct {
	name  GlfwWindowHint
	value interface{}
}

func WindowHint(hintName GlfwWindowHint, value interface{}) GlfwBuilderOption {
	return func(b GlfwBuilder) error {
		b.PushWindowHint(hintName, value)
		return nil
	}
}

func WindowTitle(title string) GlfwBuilderOption {
	return func(b GlfwBuilder) error {
		b.WindowTitle(title)
		return nil
	}
}
