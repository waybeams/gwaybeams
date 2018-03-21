package display

type LayoutHandler func(d Displayable)

type Layout int

// This pattern is probably not the way to go, but I'm having trouble finding a
// reasonable alternative. The problem here is that Layout types will not be
// user-extensible. Component definitions will only be able to refer to the
// Layouts that have been enumerated here. The benefit is that Opts objects
// will remain serializable and simply be a bag of scalars. I'm definitely
// open to suggestions here.
const (
	FlowLayout = iota
	RowLayout
	StackLayout
)

type Direction int

const (
	Horizontal = iota
	Vertical
)

func DirectionalDelegate(d Direction) func(d Displayable) {
	return func(d Displayable) {
	}
}

/*
func LayoutStack(d Displayable) {
}

func LayoutFlow(d Displayable) {
}

func LayoutRow(d Displayable) {
}
*/
