package layout

import (
	"math"
	"ui"
)

// These entities are stateless bags of hooks that allow us to apply
// the exact same layout rules on both supported axes.
var hDelegate *horizontalDelegate
var vDelegate *verticalDelegate

// Instantiate each delegate once the declarations are ready
// TODO(lbayes): These have been made stateless, refactor into a module?
func init() {
	hDelegate = &horizontalDelegate{}
	vDelegate = &verticalDelegate{}
}

func NoLayout(d ui.Displayable) {
	// noop
}

// StackLayout arranges children in a vertical flow and use stack for
// horizontal rules.
func StackLayout(d ui.Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	if vDelegate.Fixed(d) == -1 && hDelegate.Flex(d) == -1 {
		hChildrenSize := stackGetChildrenSize(hDelegate, d) + hDelegate.Padding(d)
		hDelegate.SetActualSize(d, hChildrenSize)
	}

	if vDelegate.Fixed(d) == -1 && vDelegate.Flex(d) == -1 {
		vChildrenSize := stackGetChildrenSize(vDelegate, d) + vDelegate.Padding(d)
		vDelegate.SetActualSize(d, vChildrenSize)
	}

	stackScaleChildren(hDelegate, d)
	stackScaleChildren(vDelegate, d)

	stackPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func HorizontalFlowLayout(d ui.Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	flexibleChildren := getFlexibleChildren(hDelegate, d)
	flowScaleChildren(hDelegate, d, flexibleChildren)
	stackScaleChildren(vDelegate, d)

	flowPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func VerticalFlowLayout(d ui.Displayable) {
	if d.ChildCount() == 0 {
		return
	}

	stackScaleChildren(hDelegate, d)
	flexibleChildren := getFlexibleChildren(vDelegate, d)
	flowScaleChildren(vDelegate, d, flexibleChildren)

	stackPositionChildren(hDelegate, d)
	flowPositionChildren(vDelegate, d)
}

func stackGetChildrenSize(delegate LayoutDelegate, d ui.Displayable) float64 {
	max := 0.0
	for _, child := range d.Children() {
		max = math.Max(max, delegate.Size(child))
	}
	return max
}

func notExcludedFromLayout(d ui.Displayable) bool {
	return !d.ExcludeFromLayout()
}

// Collect the layoutable children of a Displayable
func getLayoutableChildren(d ui.Displayable) []ui.Displayable {
	return d.GetFilteredChildren(notExcludedFromLayout)
}

func getFlexibleChildren(delegate LayoutDelegate, d ui.Displayable) []ui.Displayable {
	return d.GetFilteredChildren(func(child ui.Displayable) bool {
		isExcluded := child.ExcludeFromLayout()
		isFlexible := delegate.IsFlexible(child)
		return isFlexible && !isExcluded
	})
}

func getNotExcludedFromLayoutChildren(delegate LayoutDelegate, d ui.Displayable) []ui.Displayable {
	return d.GetFilteredChildren(func(child ui.Displayable) bool {
		return notExcludedFromLayout(child)
	})
}

func childIsFlexible(delegate LayoutDelegate, child ui.Displayable, flexibleChildren []ui.Displayable) bool {
	if !delegate.IsFlexible(child) {
		// The child itself does not have a flex property
		return false
	}
	// The child may have been removed from this collection because it hit a boundary dimension.
	for _, fChild := range flexibleChildren {
		if fChild == child {
			return true
		}
	}
	// The child reported a flex property, but was removed from the flex children collection,
	// most likely because it hit a boundary condition.
	return false
}

func getStaticChildren(delegate LayoutDelegate, d ui.Displayable, flexibleChildren []ui.Displayable) []ui.Displayable {
	staticChildren := d.GetFilteredChildren(func(child ui.Displayable) bool {
		return notExcludedFromLayout(child) && !childIsFlexible(delegate, child, flexibleChildren)
	})
	return staticChildren
}

func flowScaleChildren(delegate LayoutDelegate, d ui.Displayable, flexibleChildren []ui.Displayable) {
	if len(flexibleChildren) > 0 {

		unitSize, remainder := flowGetUnitSize(delegate, d, flexibleChildren)
		for index, child := range flexibleChildren {
			value := math.Floor(delegate.Flex(child) * unitSize)
			delegate.SetActualSize(child, value)

			if delegate.Size(child) < value {
				// We bumped into a size boundary, remove the limited entry and attempt to spread
				// the difference.
				flexibleChildren := append(flexibleChildren[:index], flexibleChildren[index+1:]...)
				flowScaleChildren(delegate, d, flexibleChildren)
				break
			}
			// TODO(lbayes): Break out if child failed to take the requested size
			// Consider updating the setter api to return the value that was set?
		}
		flowSpreadRemainder(delegate, flexibleChildren, remainder)
	}
}

// Position the scaled children and return the new parent dimension.
func flowPositionChildren(delegate LayoutDelegate, d ui.Displayable) {
	children := getNotExcludedFromLayoutChildren(delegate, d)
	position := delegate.PaddingFirst(d)
	gutter := d.Gutter()
	for _, child := range children {
		delegate.SetPosition(child, position)
		position = position + delegate.Size(child) + gutter
	}
	size := position + delegate.PaddingLast(d)
	if delegate.Size(d) < size {
		delegate.SetActualSize(d, size)
	}
}

func flowSpreadRemainder(delegate LayoutDelegate, flexibleChildren []ui.Displayable, remainder int) {
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
		size := delegate.Size(child)
		delegate.SetActualSize(child, size+unit)
		remainder--
	}
}

func flowGetUnitSize(delegate LayoutDelegate, d ui.Displayable, flexibleChildren []ui.Displayable) (unitSize float64, remainder int) {
	availablePixels := flowGetAvailablePixels(delegate, d, flexibleChildren)
	flexSum := flowGetFlexSum(delegate, flexibleChildren)
	if flexSum > 0.0 {
		unitSize = availablePixels / flexSum
		return availablePixels / flexSum, int(availablePixels) % int(flexSum)
	}
	return 0.0, 0
}

func flowGetFlexSum(delegate LayoutDelegate, flexibleChildren []ui.Displayable) float64 {
	sum := 0.0
	for _, child := range flexibleChildren {
		sum += delegate.Flex(child)
	}
	return math.Floor(sum)
}

func stackScaleChildren(delegate LayoutDelegate, d ui.Displayable) {
	maxChildSize := 0.0
	flexChildren := getFlexibleChildren(delegate, d)
	availablePixels := getAvailablePixels(delegate, d)

	for _, child := range flexChildren {
		delegate.SetActualSize(child, availablePixels)
		maxChildSize = math.Max(maxChildSize, delegate.Size(child))
	}

	if maxChildSize > availablePixels {
		delegate.SetActualSize(d, maxChildSize+delegate.Padding(d))
	}
}

// Get the (Size - Padding) on delegated axis for STACK layouts.
func getAvailablePixels(delegate LayoutDelegate, d ui.Displayable) float64 {
	return delegate.Size(d) - delegate.Padding(d)
}

func flowGetAvailablePixels(delegate LayoutDelegate, d ui.Displayable, flexibleChildren []ui.Displayable) float64 {
	staticChildren := getStaticChildren(delegate, d, flexibleChildren)
	staticChildrenSize := 0.0
	for _, child := range staticChildren {
		staticChildrenSize += math.Max(0.0, delegate.Size(child))
	}
	return delegate.Size(d) - delegate.Padding(d) - staticChildrenSize
}

func stackPositionChildren(delegate LayoutDelegate, d ui.Displayable) {
	align := delegate.Align(d)
	switch align {
	case ui.AlignLeft:
		fallthrough
	case ui.AlignTop:
		stackPositionChildrenFirst(delegate, d)
	case ui.AlignCenter:
		stackPositionChildrenCenter(delegate, d)
	default:
		stackPositionChildrenLast(delegate, d)
	}
}

func stackPositionChildrenFirst(delegate LayoutDelegate, d ui.Displayable) {
	// Position all children in upper left of container
	pos := delegate.PaddingFirst(d)
	for _, child := range getLayoutableChildren(d) {
		delegate.SetPosition(child, pos)
	}
}

func stackPositionChildrenCenter(delegate LayoutDelegate, d ui.Displayable) {
	// Position all children in upper left of container

	space := delegate.Size(d) - delegate.Padding(d)
	paddingFirst := delegate.PaddingFirst(d)

	for _, child := range getLayoutableChildren(d) {
		childSize := delegate.Size(child)
		pos := paddingFirst + ((space - childSize) / 2)
		delegate.SetPosition(child, pos)
	}
}

func stackPositionChildrenLast(delegate LayoutDelegate, d ui.Displayable) {
	last := delegate.Size(d) - delegate.PaddingLast(d)
	for _, child := range getLayoutableChildren(d) {
		pos := last - delegate.Size(child)
		delegate.SetPosition(child, pos)
	}
}
