package display

type GlfwBuilderOption func(g GlfwBuilder) error
type BuilderOption func(b Builder) error

func FrameRate(fps int) GlfwBuilderOption {
	return func(g GlfwBuilder) error {
		g.FrameRate(fps)
		return nil
	}
}

func WindowSize(width int, height int) GlfwBuilderOption {
	return func(g GlfwBuilder) error {
		g.WindowSize(width, height)
		return nil
	}
}

// WindowHints are how we configure GLFW windows
type windowHint struct {
	name  GlfwWindowHint
	value interface{}
}

func WindowHint(hintName GlfwWindowHint, value interface{}) GlfwBuilderOption {
	return func(g GlfwBuilder) error {
		g.PushWindowHint(hintName, value)
		return nil
	}
}

func WindowTitle(title string) GlfwBuilderOption {
	return func(g GlfwBuilder) error {
		g.WindowTitle(title)
		return nil
	}
}

func BuildAndLoop(composer ComponentComposer) GlfwBuilderOption {
	return func(g GlfwBuilder) error {
		g.Build(composer)
		g.Loop()
		return nil
	}
}
