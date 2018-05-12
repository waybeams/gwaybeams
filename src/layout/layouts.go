package layout

import (
	"fmt"
	"math"
	"spec"
	"surface"
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

func selectLayout(r spec.Reader) spec.LayoutHandler {
	switch r.LayoutType() {
	case spec.NoLayoutType:
		return None
	case spec.StackLayoutType:
		return Stack
	case spec.HorizontalFlowLayoutType:
		return FlowHorizontal
	case spec.VerticalFlowLayoutType:
		return FlowVertical
	default:
		panic(fmt.Sprintf("ERROR: Requested LayoutTypeValue (%v) is not supported:", r.LayoutType()))
		return nil
	}
}

func execLayout(r spec.ReadWriter) {
	selectLayout(r)(r)
}

func Measure(r spec.ReadWriter, s spec.Surface) {
	// Leaf first traversal
	for _, child := range r.Children() {
		Measure(child, s)
	}
	if r.IsMeasured() {
		// Ask the Spec to measure itself, which will update the size or
		// size boundaries depending on the implementation.
		r.Measure(s)
	}
}

// Layout the provided control and all of it's children.
func Layout(r spec.ReadWriter, s spec.Surface) spec.ReadWriter {
	s = surface.NewOffsetSurface(r, s)
	Measure(r, s)
	execLayout(r)
	layoutChildren(r)
	return r
}

func layoutChildren(r spec.ReadWriter) {
	// NOTE(lbayes): Do not recurse back to entry point, before
	// getting into here, we perform a depth first Measure from
	// the Layout function.
	for _, child := range r.Children() {
		execLayout(child)
		layoutChildren(child)
	}
}

func None(d spec.ReadWriter) {
	// noop
}

// StackLayout arranges children in a vertical flow and use stack for
// horizontal rules.
func Stack(d spec.ReadWriter) {
	if d.ChildCount() == 0 {
		return
	}

	if vDelegate.Fixed(d) == 0 && hDelegate.Flex(d) == 0 {
		hChildrenSize := stackGetChildrenSize(hDelegate, d) + hDelegate.Padding(d)
		hDelegate.SetActualSize(d, hChildrenSize)
	}

	if vDelegate.Fixed(d) == 0 && vDelegate.Flex(d) == 0 {
		vChildrenSize := stackGetChildrenSize(vDelegate, d) + vDelegate.Padding(d)
		vDelegate.SetActualSize(d, vChildrenSize)
	}

	stackScaleChildren(hDelegate, d)
	stackScaleChildren(vDelegate, d)

	stackPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func FlowHorizontal(d spec.ReadWriter) {
	if d.ChildCount() == 0 {
		return
	}

	flexibleChildren := getFlexibleChildren(hDelegate, d)
	flowScaleChildren(hDelegate, d, flexibleChildren)
	stackScaleChildren(vDelegate, d)

	flowPositionChildren(hDelegate, d)
	stackPositionChildren(vDelegate, d)
}

func FlowVertical(d spec.ReadWriter) {
	if d.ChildCount() == 0 {
		return
	}

	stackScaleChildren(hDelegate, d)
	flexibleChildren := getFlexibleChildren(vDelegate, d)
	flowScaleChildren(vDelegate, d, flexibleChildren)

	stackPositionChildren(hDelegate, d)
	flowPositionChildren(vDelegate, d)
}

func stackGetChildrenSize(delegate Delegate, d spec.ReadWriter) float64 {
	max := 0.0
	for _, child := range d.Children() {
		max = math.Max(max, delegate.Size(child))
	}
	return max
}

func notExcludedFromLayout(d spec.Reader) bool {
	return !d.ExcludeFromLayout()
}

// Collect the layoutable children of a Displayable
func getLayoutableChildren(d spec.ReadWriter) []spec.ReadWriter {
	return spec.FilteredChildren(d, notExcludedFromLayout)
}

func getFlexibleChildren(delegate Delegate, d spec.ReadWriter) []spec.ReadWriter {
	return spec.FilteredChildren(d, func(child spec.Reader) bool {
		isExcluded := child.ExcludeFromLayout()
		isFlexible := delegate.IsFlexible(child)
		return isFlexible && !isExcluded
	})
}

func getNotExcludedFromLayoutChildren(delegate Delegate, d spec.ReadWriter) []spec.ReadWriter {
	return spec.FilteredChildren(d, func(child spec.Reader) bool {
		return notExcludedFromLayout(child)
	})
}

func childIsFlexible(delegate Delegate, child spec.Reader, flexibleChildren []spec.ReadWriter) bool {
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

func getStaticChildren(delegate Delegate, d spec.ReadWriter, flexibleChildren []spec.ReadWriter) []spec.ReadWriter {
	staticChildren := spec.FilteredChildren(d, func(child spec.Reader) bool {
		return notExcludedFromLayout(child) && !childIsFlexible(delegate, child, flexibleChildren)
	})
	return staticChildren
}

func flowScaleChildren(delegate Delegate, d spec.ReadWriter, flexibleChildren []spec.ReadWriter) {
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
func flowPositionChildren(delegate Delegate, d spec.ReadWriter) {
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

func flowSpreadRemainder(delegate Delegate, flexibleChildren []spec.ReadWriter, remainder float64) {
	count := float64(len(flexibleChildren))
	for _, child := range flexibleChildren {
		var unit float64
		if remainder <= count {
			unit = 1
		} else {
			unit = remainder / count
		}
		if remainder == 0 {
			return
		}
		size := delegate.Size(child)
		delegate.SetActualSize(child, size+unit)
		remainder--
	}
}

func flowGetUnitSize(delegate Delegate, d spec.ReadWriter, flexibleChildren []spec.ReadWriter) (unitSize float64, remainder float64) {
	availablePixels := flowGetAvailablePixels(delegate, d, flexibleChildren)
	flexSum := flowGetFlexSum(delegate, flexibleChildren)
	if flexSum > 0 {
		unitSize = math.Floor(availablePixels / flexSum)
		return unitSize, math.Mod(availablePixels, flexSum)
	}
	return 0, 0
}

func flowGetFlexSum(delegate Delegate, flexibleChildren []spec.ReadWriter) float64 {
	sum := 0.0
	for _, child := range flexibleChildren {
		sum += delegate.Flex(child)
	}
	return math.Floor(sum)
}

func stackScaleChildren(delegate Delegate, d spec.ReadWriter) {
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
func getAvailablePixels(delegate Delegate, d spec.ReadWriter) float64 {
	return delegate.Size(d) - delegate.Padding(d)
}

func flowGetAvailablePixels(delegate Delegate, d spec.ReadWriter, flexibleChildren []spec.ReadWriter) float64 {
	staticChildren := getStaticChildren(delegate, d, flexibleChildren)
	staticChildrenSize := 0.0
	for _, child := range staticChildren {
		staticChildrenSize += math.Max(0.0, delegate.Size(child))
	}
	return delegate.Size(d) - delegate.Padding(d) - staticChildrenSize
}

func stackPositionChildren(delegate Delegate, d spec.ReadWriter) {
	align := delegate.Align(d)
	switch align {
	case spec.AlignLeft:
		fallthrough
	case spec.AlignTop:
		stackPositionChildrenFirst(delegate, d)
	case spec.AlignCenter:
		stackPositionChildrenCenter(delegate, d)
	default:
		stackPositionChildrenLast(delegate, d)
	}
}

func stackPositionChildrenFirst(delegate Delegate, d spec.ReadWriter) {
	// Position all children in upper left of container
	pos := delegate.PaddingFirst(d)
	for _, child := range getLayoutableChildren(d) {
		delegate.SetPosition(child, pos)
	}
}

func stackPositionChildrenCenter(delegate Delegate, d spec.ReadWriter) {
	// Position all children in upper left of container

	space := delegate.Size(d) - delegate.Padding(d)
	paddingFirst := delegate.PaddingFirst(d)

	for _, child := range getLayoutableChildren(d) {
		childSize := delegate.Size(child)
		pos := paddingFirst + ((space - childSize) / 2)
		delegate.SetPosition(child, pos)
	}
}

func stackPositionChildrenLast(delegate Delegate, d spec.ReadWriter) {
	last := delegate.Size(d) - delegate.PaddingLast(d)
	for _, child := range getLayoutableChildren(d) {
		pos := last - delegate.Size(child)
		delegate.SetPosition(child, pos)
	}
}
