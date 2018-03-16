package example

// Semantic exploration for component creation

// Library defined components
func Window(args ...interface{}) {}
func Box(args ...interface{})    {}
func VBox(args ...interface{})   {}
func HBox(args ...interface{})   {}
func Header(args ...interface{}) {}

type Opts struct {
	FlexHeight int
	FlexWidth  int
	Height     int
	Width      int
}
