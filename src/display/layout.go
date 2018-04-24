package display

import (
	"math"
)

type LayoutAxis int

const (
	LayoutHorizontal = iota
	LayoutVertical
)

// LayoutTypeValue is a serializable enum for selecting a layout scheme.
// This pattern is probably not the way to go, but I'm having trouble finding a
// reasonable alternative. The problem here is that LayoutHandler types will not be
// user-extensible. Box definitions will only be able to refer to the
// Layouts that have been enumerated here. The benefit is that ComponentModel objects
// will remain serializable and simply be a bag of scalars. I'm definitely
// open to suggestions.
type LayoutTypeValue int

const (
	StackLayoutType = iota
	VerticalFlowLayoutType
	HorizontalFlowLayoutType
	RowLayoutType
)

// Alignment is used represent alignment of Component children, text or any other
// alignable entities.
type Alignment int

const (
	AlignBottom = iota
	AlignLeft
	AlignRight
	AlignTop
	AlignCenter
)

// LayoutHandler is a concrete implementation of a given layout. These handlers
// are pure functions that accept a Displayable and manage the scale and
// position of the children for that element.
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

// StackLayout arranges children in a vertical flow and use stack for
// horizontal rules.
func StackLayout(d Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	if hDelegate.GetFixed(d) == -1 && hDelegate.GetFlex(d) == -1 {
		hDelegate.ActualSize(d, hDelegate.GetChildrenSize(d))
	}

	if vDelegate.GetFixed(d) == -1 && vDelegate.GetFlex(d) == -1 {
		vDelegate.ActualSize(d, vDelegate.GetChildrenSize(d))
	}

	stackScaleChildren(hDelegate, d)
	stackScaleChildren(vDelegate, d)

	stackPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func HorizontalFlowLayout(d Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	flowScaleChildren(hDelegate, d)
	stackScaleChildren(vDelegate, d)

	flowPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func VerticalFlowLayout(d Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	stackScaleChildren(hDelegate, d)
	flowScaleChildren(vDelegate, d)

	stackPositionChildren(hDelegate, d)
	flowPositionChildren(vDelegate, d)
}

func notExcludedFromLayout(d Displayable) bool {
	return !d.ExcludeFromLayout()
}

// Collect the layoutable children of a Displayable
func getLayoutableChildren(d Displayable) []Displayable {
	return d.GetFilteredChildren(notExcludedFromLayout)
}

func getFlexibleChildren(delegate LayoutDelegate, d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		isExcluded := child.ExcludeFromLayout()
		isFlexible := delegate.GetIsFlexible(child)
		return isFlexible && !isExcluded
	})
}

func getNotExcludedFromLayoutChildren(delegate LayoutDelegate, d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child)
	})
}

func getStaticChildren(delegate LayoutDelegate, d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child) && !delegate.GetIsFlexible(child)
	})
}

func getStaticSize(delegate LayoutDelegate, d Displayable) float64 {
	sum := 0.0
	staticChildren := getStaticChildren(delegate, d)
	for _, child := range staticChildren {
		sum += math.Max(0.0, delegate.GetSize(child))
	}
	return sum
}

func flowScaleChildren(delegate LayoutDelegate, d Displayable) {
	flexibleChildren := getFlexibleChildren(delegate, d)

	if len(flexibleChildren) > 0 {

		unitSize, remainder := flowGetUnitSize(delegate, d, flexibleChildren)
		for _, child := range flexibleChildren {
			value := math.Floor(delegate.GetFlex(child) * unitSize)
			delegate.ActualSize(child, value)
			// TODO(lbayes): Break out if child failed to take the requested size
			// Consider updating the setter api to return the value that was set?
		}
		flowSpreadRemainder(delegate, flexibleChildren, remainder)
	}
}

// Position the scaled children and return the new parent dimension.
func flowPositionChildren(delegate LayoutDelegate, d Displayable) {
	children := getNotExcludedFromLayoutChildren(delegate, d)
	position := delegate.GetPaddingFirst(d)
	gutter := d.Gutter()
	for _, child := range children {
		delegate.Position(child, position)
		position = position + delegate.GetSize(child) + gutter
	}
	size := position + delegate.GetPaddingLast(d)
	if delegate.GetSize(d) < size {
		delegate.ActualSize(d, size)
	}
}

func flowSpreadRemainder(delegate LayoutDelegate, flexibleChildren []Displayable, remainder int) {
	count := len(flexibleChildren)
	for _, child := range flexibleChildren {
		var unit float64
		if remainder <= count {
			unit = 1
		} else {
			unit = float64(remainder / count)
		}
		if remainder == 0 {
			return
		}
		size := delegate.GetSize(child)
		delegate.ActualSize(child, size+unit)
		remainder--
	}
}

func flowGetUnitSize(delegate LayoutDelegate, d Displayable, flexibleChildren []Displayable) (unitSize float64, remainder int) {
	availablePixels := flowGetAvailablePixels(delegate, d)
	flexSum := flowGetFlexSum(delegate, flexibleChildren)
	if flexSum > 0.0 {
		unitSize = availablePixels / flexSum
		return availablePixels / flexSum, int(availablePixels) % int(flexSum)
	}
	return 0.0, 0
}

func flowGetFlexSum(delegate LayoutDelegate, flexibleChildren []Displayable) float64 {
	sum := 0.0
	for _, child := range flexibleChildren {
		sum += delegate.GetFlex(child)
	}
	return math.Floor(sum)
}

func stackScaleChildren(delegate LayoutDelegate, d Displayable) {
	flexChildren := getFlexibleChildren(delegate, d)

	if len(flexChildren) == 0 {
		return
	}

	availablePixels := getAvailablePixels(delegate, d)

	for _, child := range flexChildren {
		delegate.ActualSize(child, availablePixels)
	}
}

// Get the (Size - Padding) on delegated axis for STACK layouts.
func getAvailablePixels(delegate LayoutDelegate, d Displayable) float64 {
	return delegate.GetSize(d) - delegate.GetPadding(d)
}

func flowGetAvailablePixels(delegate LayoutDelegate, d Displayable) float64 {
	staticChildren := getStaticChildren(delegate, d)
	staticChildrenSize := 0.0
	for _, child := range staticChildren {
		staticChildrenSize += math.Max(0.0, delegate.GetSize(child))
	}
	return delegate.GetSize(d) - delegate.GetPadding(d) - staticChildrenSize
}

func stackPositionChildren(delegate LayoutDelegate, d Displayable) {
	// TODO(lbayes): Work with alignment (first, center, last == left, center, right or top, center, bottom)

	maxChildSize := 0.0
	// Position all children in upper left of container
	pos := delegate.GetPaddingFirst(d)
	for _, child := range getLayoutableChildren(d) {
		delegate.Position(child, pos)
		maxChildSize = math.Max(maxChildSize, delegate.GetSize(child))
	}

	size := maxChildSize + delegate.GetPadding(d)
	if delegate.GetSize(d) < size {
		delegate.ActualSize(d, size)
	}
}
