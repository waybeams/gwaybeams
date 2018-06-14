package nano

import "github.com/shibukawa/nanovgo"

type Option func(s *Surface)

func Width(width float64) Option {
	return func(s *Surface) {
		s.SetWidth(width)
	}
}

func Height(height float64) Option {
	return func(s *Surface) {
		s.SetHeight(height)
	}
}

func AntiAlias() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.AntiAlias)
	}
}

func StencilStrokes() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.StencilStrokes)
	}
}

func Debug() Option {
	return func(s *Surface) {
		s.flags = append(s.flags, nanovgo.Debug)
	}
}

func AddFont(name, path string) Option {
	return func(s *Surface) {
		s.AddFont(name, path)
	}
}
