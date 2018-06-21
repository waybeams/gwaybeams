package browser

type SurfaceOption func(s *Surface)

func AddFont(name, path string) SurfaceOption {
	return func(s *Surface) {
		s.AddFont(name, path)
	}
}
