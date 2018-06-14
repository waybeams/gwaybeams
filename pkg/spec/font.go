package spec

// Font is the primitive entity that is used by measurable components to
// determine text dimensions at runtime.
type Font interface {
	Bounds(value string) (width float32, bounds []float32)
	IsCreated() bool
	Name() string
	OnCreated()
	Path() string
	SetAlign(align Alignment)
	SetSize(size float32)
	VerticalMetrics() (ascender, descender, lineHeight float32)
}
