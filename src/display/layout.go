package display

import "fmt"

type LayoutHandler func(d Displayable)

// These entities are stateless bags of hooks that allow us to apply
// the exact same layout rules on both supported axes.
var hDelegate *horizontalDelegate
var vDelegate *verticalDelegate

// Instantiate each delegate once the declarations are ready
func init() {
	hDelegate = &horizontalDelegate{}
	vDelegate = &verticalDelegate{}
}

func notExcludedFromLayout(d Displayable) bool {
	return !d.GetExcludeFromLayout()
}

func isFlexible(d Displayable) bool {
	return d.GetFlexWidth() > 0 || d.GetFlexHeight() > 0
}

// Collect the layoutable children of a Displayable
func GetLayoutableChildren(d Displayable) []Displayable {
	return d.GetFilteredChildren(notExcludedFromLayout)
}

func GetFlexibleChildren(d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child) && isFlexible(child)
	})
}

func GetStaticChildren(d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child) && !isFlexible(child)
	})
}

func DirectionalDelegate(d Direction) func(d Displayable) {
	var delegate LayoutDelegate
	switch d {
	case Horizontal:
		delegate = hDelegate
	case Vertical:
		delegate = vDelegate
	}
	return func(d Displayable) {
		fmt.Println("delegate", delegate.GetFixed(d))
	}
}

type horizontalDelegate struct {
}

func (h *horizontalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedWidth()
}

type verticalDelegate struct {
}

func (h *verticalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedHeight()
}

/*
func LayoutStack(d Displayable) {
}

func LayoutFlow(d Displayable) {
}

func LayoutRow(d Displayable) {
}
*/

type LayoutDelegate interface {
	GetFixed(d Displayable) float64
}
