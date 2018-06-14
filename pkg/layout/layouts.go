package layout

import (
	"math"

	"github.com/waybeams/waybeams/pkg/spec"
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

// Measure the provided tree, using leaf-first traversal.
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
	s = spec.NewOffsetSurface(r, s)
	Measure(r, s)
	if r.ChildCount() == 0 {
		return r
	}
	w := hDelegate.LayoutSpec(r)
	h := vDelegate.LayoutSpec(r)
	r.SetChildrenWidth(w)
	r.SetChildrenHeight(h)
	return r
}

// None Layout will prevent any automated layout from the current node through all children.
func None(delegate Delegate, d spec.ReadWriter) (childrenSize float64) {
	return delegate.Size(d)
}

func layoutStackChildren(d spec.ReadWriter, delegate Delegate) float64 {
	maxSize := 0.0
	for _, child := range d.Children() {
		maxSize = math.Max(maxSize, delegate.LayoutSpec(child))
	}
	return maxSize
}

func layoutFlowChildren(delegate Delegate, d spec.ReadWriter) (childrenSize float64) {
	var lastChild spec.ReadWriter
	for _, lastChild = range d.Children() {
		delegate.LayoutSpec(lastChild)
	}
	if lastChild == nil {
		return 0
	}
	return delegate.Size(lastChild) + delegate.Position(lastChild)
}

// StackOnAxis performs a Stack layout on the provided delegate axis.
func StackOnAxis(delegate Delegate, d spec.ReadWriter) (childrenSize float64) {
	if d.ChildCount() == 0 {
		return delegate.Size(d)
	}

	// Update the children size based on first attempt at scaling immediate
	// children.
	childrenSize = stackScaleChildren(delegate, d)
	delegate.SetChildrenSize(d, childrenSize)

	// Update the children size based on deeper traversal into layouts beyond
	// immediate children.
	childrenSize = layoutStackChildren(d, delegate)
	delegate.SetChildrenSize(d, childrenSize)

	stackPositionChildren(delegate, d)
	return childrenSize
}

// FlowOnAxis performs a Flow layout on the provided delegate axis.
func FlowOnAxis(delegate Delegate, d spec.ReadWriter) (childrenSize float64) {
	if d.ChildCount() == 0 {
		return delegate.Size(d)
	}
	flowScaleChildren(delegate, d, nil)

	childrenSize = layoutFlowChildren(delegate, d)
	delegate.SetChildrenSize(d, childrenSize)

	childrenSize = flowPositionChildren(delegate, d)
	delegate.SetChildrenSize(d, childrenSize)
	return childrenSize
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

func getNotExcludedFromLayoutChildren(d spec.ReadWriter) []spec.ReadWriter {
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
	if flexibleChildren == nil {
		flexibleChildren = getFlexibleChildren(delegate, d)
	}

	if len(flexibleChildren) > 0 {
		unitSize, remainder := flowGetUnitSize(delegate, d, flexibleChildren)
		for index, child := range flexibleChildren {
			requestedSize := math.Floor(delegate.Flex(child) * unitSize)
			actualSize := delegate.SetSize(child, requestedSize)

			if actualSize != requestedSize {
				// We bumped into a size boundary, remove the size-limited
				// entry and restart in an attempt to spread the difference.
				flexibleChildren := append(flexibleChildren[:index], flexibleChildren[index+1:]...)
				flowScaleChildren(delegate, d, flexibleChildren)
				return
			}
		}
		flowSpreadRemainder(delegate, flexibleChildren, remainder)
	}
}

// Position the scaled children and return the new parent dimension.
func flowPositionChildren(delegate Delegate, s spec.ReadWriter) (childrenSize float64) {
	children := getNotExcludedFromLayoutChildren(s)
	paddingFirst := delegate.PaddingFirst(s)
	position := paddingFirst
	gutter := s.Gutter()
	for _, child := range children {
		delegate.SetPosition(child, position)
		position = position + delegate.Size(child) + gutter
	}
	return position - gutter - paddingFirst
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
		delegate.SetSize(child, size+unit)
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

func stackScaleChildren(delegate Delegate, d spec.ReadWriter) (childrenSize float64) {
	childrenSize = 0.0
	children := getLayoutableChildren(d)
	flexibleChildren := getFlexibleChildren(delegate, d)
	availablePixels := getAvailablePixels(delegate, d)

	for _, child := range children {
		if childIsFlexible(delegate, child, flexibleChildren) {
			delegate.SetSize(child, availablePixels)
		}
		childrenSize = math.Max(childrenSize, delegate.Size(child))
	}
	return childrenSize
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
		// case spec.AlignRight:
		// fallthrough
		// case spec.AlignBottom:
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
