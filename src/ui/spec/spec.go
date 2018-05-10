package spec

type Option func(w ReadWriter)

type Writer interface {
	SetChildren(children []ReadWriter)
	SetComposer(composer interface{})
	SetHeight(heigth float32)
	SetKey(key string)
	SetName(name string)
	SetParent(parent ReadWriter)
	SetWidth(width float32)
	SetX(x float32)
	SetY(y float32)
}

type Reader interface {
	ChildAt(index int) ReadWriter
	ChildCount() int
	Children() []ReadWriter
	Composer() interface{}
	Height() float32
	Name() string
	Parent() ReadWriter
	Key() string
	Width() float32
	X() float32
	Y() float32
}

type ReadWriter interface {
	Reader
	Writer
}

// ControlSpec is an internallly configurable, externally read-only bag of
// state that describes a user interface element.
type ControlSpec struct {
	children []ReadWriter
	composer interface{}
	height   float32
	key      string
	name     string
	parent   ReadWriter
	width    float32
	x        float32
	y        float32
}

func (c *ControlSpec) ChildAt(index int) ReadWriter {
	// if c.children != nil && len(c.children) > index-1 {
	return c.children[index]
	// }
	// return nil
}

func (c *ControlSpec) ChildCount() int {
	return len(c.children)
}

func (c *ControlSpec) Children() []ReadWriter {
	return c.children
}

func (c *ControlSpec) Composer() interface{} {
	return c.composer
}

func (c *ControlSpec) Height() float32 {
	return c.height
}

func (c *ControlSpec) Key() string {
	return c.key
}

func (c *ControlSpec) Name() string {
	return c.name
}

func (c *ControlSpec) Parent() ReadWriter {
	return c.parent
}

func (c *ControlSpec) Width() float32 {
	return c.width
}

func (c *ControlSpec) X() float32 {
	return c.x
}

func (c *ControlSpec) Y() float32 {
	return c.x
}

func (c *ControlSpec) SetChildren(children []ReadWriter) {
	c.children = children
}

func (c *ControlSpec) SetComposer(composer interface{}) {
	c.composer = composer
}

func (c *ControlSpec) SetHeight(height float32) {
	c.height = height
}

func (c *ControlSpec) SetKey(key string) {
	c.key = key
}

func (c *ControlSpec) SetName(name string) {
	c.name = name
}

func (c *ControlSpec) SetParent(parent ReadWriter) {
	c.parent = parent
}

func (c *ControlSpec) SetWidth(width float32) {
	c.width = width
}

func (c *ControlSpec) SetX(x float32) {
	c.x = x
}

func (c *ControlSpec) SetY(y float32) {
	c.y = y
}

// Spec Options require a spec.ReadWriter and will apply values to the provided
// instance.

func Bag(options ...Option) Option {
	return func(rw ReadWriter) {
		Apply(rw, options...)
	}
}

func Child(child ReadWriter) Option {
	return func(rw ReadWriter) {
		rw.SetChildren(append(rw.Children(), child))
		child.SetParent(rw)
	}
}

func Height(value float32) Option {
	return func(rw ReadWriter) {
		rw.SetHeight(value)
	}
}

func Key(value string) Option {
	return func(rw ReadWriter) {
		rw.SetKey(value)
	}
}

func Name(name string) Option {
	return func(rw ReadWriter) {
		rw.SetName(name)
	}
}

func Width(value float32) Option {
	return func(rw ReadWriter) {
		rw.SetWidth(value)
	}
}

func X(value float32) Option {
	return func(rw ReadWriter) {
		rw.SetX(value)
	}
}

func Y(value float32) Option {
	return func(rw ReadWriter) {
		rw.SetY(value)
	}
}

// ApplyAll will take arbitrary slices of Options and will apply each of them
// in order from left to right.
func ApplyAll(rw ReadWriter, optionSets ...[]Option) ReadWriter {
	options := []Option{}
	for _, optionSet := range optionSets {
		options = append(options, optionSet...)
	}
	return Apply(rw, options...)
}

// Apply will call each provided Option with the provided ReadWriter.
func Apply(rw ReadWriter, options ...Option) ReadWriter {
	for _, option := range options {
		option(rw)
	}
	return rw
}
