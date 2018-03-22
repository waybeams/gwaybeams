package display

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

func GetFlexibleChildren(delegate LayoutDelegate, d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child) && delegate.GetIsFlexible(child)
	})
}

func GetStaticChildren(d Displayable) []Displayable {
	return d.GetFilteredChildren(func(child Displayable) bool {
		return notExcludedFromLayout(child) && !isFlexible(child)
	})
}

func GetStaticSize(delegate LayoutDelegate, d Displayable) float64 {
	sum := 0.0
	staticChildren := GetStaticChildren(d)
	for _, child := range staticChildren {
		sum += delegate.GetSize(child)
	}
	return sum
}

// Arrange children in a vertical flow and use stack for horizontal rules.
func StackLayout(d Displayable) {
	if d.GetChildCount() == 0 {
		return
	}

	if hDelegate.GetFixed(d) == 0 && hDelegate.GetFlex(d) == 0 {
		hDelegate.ActualSize(d, hDelegate.GetChildrenSize(d))
	}

	if vDelegate.GetFixed(d) == 0 && vDelegate.GetFlex(d) == 0 {
		vDelegate.ActualSize(d, vDelegate.GetChildrenSize(d))
	}

	StackScaleChildren(hDelegate, d)
	StackScaleChildren(vDelegate, d)

	StackPositionChildren(hDelegate, d)
	StackPositionChildren(vDelegate, d)
}

func StackScaleChildren(delegate LayoutDelegate, d Displayable) {
	flexChildren := GetFlexibleChildren(delegate, d)

	if len(flexChildren) == 0 {
		return
	}

	availablePixels := StackGetAvailablePixels(delegate, d)

	for _, child := range flexChildren {
		delegate.ActualSize(child, availablePixels)
	}
}

// Get the (Size - Padding) on delegated axis for STACK layouts.
// NOTE: Flow layouts will also take into account the non-flexible children.
func StackGetAvailablePixels(delegate LayoutDelegate, d Displayable) float64 {
	return delegate.GetSize(d) - delegate.GetPadding(d)
}

func StackGetUnitSize(delegate LayoutDelegate, d Displayable, flexPixels float64) float64 {
	return delegate.GetFlex(d) * flexPixels
}

func StackPositionChildren(delegate LayoutDelegate, d Displayable) {
	// TODO(lbayes): Work with alignment (first, center, last == left, center, right or top, center, bottom)

	// Position all children in upper left of container
	pos := delegate.GetPaddingFirst(d)
	for _, child := range GetLayoutableChildren(d) {
		delegate.Position(child, pos)
	}
}

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct {
}

func (h *horizontalDelegate) ActualSize(d Displayable, size float64) {
	d.ActualWidth(size)
}

func (h *horizontalDelegate) GetActualSize(d Displayable) float64 {
	return d.GetActualWidth()
}

func (h *horizontalDelegate) GetAlign(d Displayable) Alignment {
	return d.GetHAlign()
}

func (h *horizontalDelegate) GetChildrenSize(d Displayable) float64 {
	return 0.0
}

func (h *horizontalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedWidth()
}

func (h *horizontalDelegate) GetFlex(d Displayable) float64 {
	return d.GetFlexWidth()
}

func (h *horizontalDelegate) GetIsFlexible(d Displayable) bool {
	return d.GetFlexWidth() > 0
}

func (h *horizontalDelegate) GetMinSize(d Displayable) float64 {
	return d.GetMinWidth()
}

func (h *horizontalDelegate) GetPadding(d Displayable) float64 {
	return d.GetHorizontalPadding()
}

func (h *horizontalDelegate) GetPaddingFirst(d Displayable) float64 {
	return d.GetPaddingLeft()
}

func (h *horizontalDelegate) GetPaddingLast(d Displayable) float64 {
	return d.GetPaddingRight()
}

func (h *horizontalDelegate) GetPosition(d Displayable) float64 {
	return d.GetX()
}

func (h *horizontalDelegate) GetPreferred(d Displayable) float64 {
	return d.GetPrefWidth()
}

func (h *horizontalDelegate) GetSize(d Displayable) float64 {
	return d.GetWidth()
}

func (h *horizontalDelegate) Position(d Displayable, pos float64) {
	d.X(pos)
}

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct {
}

func (v *verticalDelegate) ActualSize(d Displayable, size float64) {
	d.ActualHeight(size)
}

func (v *verticalDelegate) GetActualSize(d Displayable) float64 {
	return d.GetActualHeight()
}

func (v *verticalDelegate) GetAlign(d Displayable) Alignment {
	return d.GetVAlign()
}

func (v *verticalDelegate) GetChildrenSize(d Displayable) float64 {
	return 0.0
}

func (v *verticalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedHeight()
}

func (v *verticalDelegate) GetFlex(d Displayable) float64 {
	return d.GetFlexHeight()
}

func (v *verticalDelegate) GetIsFlexible(d Displayable) bool {
	return d.GetFlexHeight() > 0
}

func (v *verticalDelegate) GetMinSize(d Displayable) float64 {
	return d.GetMinHeight()
}

func (v *verticalDelegate) GetPadding(d Displayable) float64 {
	return d.GetVerticalPadding()
}

func (v *verticalDelegate) GetPaddingFirst(d Displayable) float64 {
	return d.GetPaddingTop()
}

func (v *verticalDelegate) GetPaddingLast(d Displayable) float64 {
	return d.GetPaddingBottom()
}

func (v *verticalDelegate) GetPosition(d Displayable) float64 {
	return d.GetY()
}

func (v *verticalDelegate) GetPreferred(d Displayable) float64 {
	return d.GetPrefHeight()
}

func (v *verticalDelegate) GetSize(d Displayable) float64 {
	return d.GetHeight()
}

func (v *verticalDelegate) GetStaticSize(d Displayable) float64 {
	return 0.0
}

func (v *verticalDelegate) Position(d Displayable, pos float64) {
	d.Y(pos)
}

type LayoutDelegate interface {
	ActualSize(d Displayable, size float64)
	GetActualSize(d Displayable) float64
	GetAlign(d Displayable) Alignment
	GetChildrenSize(d Displayable) float64
	GetFixed(d Displayable) float64
	GetFlex(d Displayable) float64 // GetPercent?
	GetIsFlexible(d Displayable) bool
	GetMinSize(d Displayable) float64
	GetPadding(d Displayable) float64
	GetPaddingFirst(d Displayable) float64
	GetPaddingLast(d Displayable) float64
	GetPosition(d Displayable) float64
	GetPreferred(d Displayable) float64
	GetSize(d Displayable) float64
	Position(d Displayable, pos float64)
}
