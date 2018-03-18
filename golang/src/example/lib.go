package example

// Semantic exploration for component creation

// Library defined components
func Window(args ...interface{}) {}
func Box(args ...interface{})    {}
func VBox(args ...interface{})   {}
func HBox(args ...interface{})   {}
func Header(args ...interface{}) {}

// Traits entries
func Traits(args ...interface{}) {}
func Trait(args ...interface{})  {}

func On(args ...interface{})     {}
func Resize(args ...interface{}) {}

// Style entries
func For(args ...interface{})    {}
func Styles(args ...interface{}) {}

// Concrete styles
func BgColor(color uint)       {}
func StrokeColor(color int)    {}
func StrokeSize(size int)      {}
func StrokeStyle(style string) {}
func FontSize(size int)        {}
func FontWeight(weight Weight) {}

type Weight int

const (
	Bold   Weight = iota
	Normal Weight = iota
	Italic Weight = iota
)

type Opts struct {
	FlexHeight int
	FlexWidth  int
	Height     int
	Width      int
}
