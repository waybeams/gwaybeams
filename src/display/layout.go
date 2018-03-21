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

// Collect the layoutable children of a Displayable
func GetLayoutableChildren(d Displayable) []Displayable {
	children := []Displayable{}
	for i := 0; i < d.GetChildCount(); i++ {
		child := d.GetChildAt(i)
		if !child.GetExcludeFromLayout() {
			children = append(children, child)
		}
	}
	return children
}

func GetFlexibleChildren(d Displayable) []Displayable {
	return []Displayable{}
}

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
