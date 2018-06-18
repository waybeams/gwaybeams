package browser

import "github.com/gopherjs/gopherjs/js"

type SurfaceOption func(s *Surface)

/*
func Width(width float64) SurfaceOption {
	return func(s *Surface) {
		// s.SetWidth(width)
	}
}

func Height(height float64) SurfaceOption {
	return func(s *Surface) {
		// s.SetHeight(height)
	}
}

func AntiAlias() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.AntiAlias)
	}
}

func StencilStrokes() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.StencilStrokes)
	}
}

func Debug() SurfaceOption {
	return func(s *Surface) {
		// s.flags = append(s.flags, nanovgo.Debug)
	}
}
*/

func Canvas(canvas *js.Object) SurfaceOption {
	return func(s *Surface) {
		s.canvas = canvas
	}
}

func SurfaceFont(name, path string) SurfaceOption {
	return func(s *Surface) {
		s.AddFont(name, path)
	}
}
