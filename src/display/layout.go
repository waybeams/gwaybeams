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

func DirectionalDelegate(d LayoutDirection) func(d Displayable) {
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

func GetStaticSize(delegate LayoutDelegate, d Displayable) float64 {
	sum := 0.0
	staticChildren := GetStaticChildren(d)
	for _, child := range staticChildren {
		sum += delegate.GetSize(child)
	}
	return sum
}

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct {
}

func (h *horizontalDelegate) GetActual(d Displayable) float64 {
	return d.GetActualWidth()
}

func (h *horizontalDelegate) GetAlign(d Displayable) Alignment {
	return d.GetHAlign()
}

func (h *horizontalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedWidth()
}

func (h *horizontalDelegate) GetFlex(d Displayable) float64 {
	return d.GetFlexWidth()
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

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct {
}

func (h *verticalDelegate) GetActual(d Displayable) float64 {
	return d.GetActualHeight()
}

func (h *verticalDelegate) GetAlign(d Displayable) Alignment {
	return d.GetVAlign()
}

func (h *verticalDelegate) GetFixed(d Displayable) float64 {
	return d.GetFixedHeight()
}

func (h *verticalDelegate) GetFlex(d Displayable) float64 {
	return d.GetFlexHeight()
}

func (h *verticalDelegate) GetMinSize(d Displayable) float64 {
	return d.GetMinHeight()
}

func (h *verticalDelegate) GetPadding(d Displayable) float64 {
	return d.GetVerticalPadding()
}

func (h *verticalDelegate) GetPaddingFirst(d Displayable) float64 {
	return d.GetPaddingTop()
}

func (h *verticalDelegate) GetPaddingLast(d Displayable) float64 {
	return d.GetPaddingBottom()
}

func (h *verticalDelegate) GetPosition(d Displayable) float64 {
	return d.GetY()
}

func (h *verticalDelegate) GetPreferred(d Displayable) float64 {
	return d.GetPrefHeight()
}

func (h *verticalDelegate) GetSize(d Displayable) float64 {
	return d.GetHeight()
}

func (h *verticalDelegate) GetStaticSize(d Displayable) float64 {
	return 0.0
}

type LayoutDelegate interface {
	GetActual(d Displayable) float64
	GetAlign(d Displayable) Alignment
	GetFixed(d Displayable) float64
	GetFlex(d Displayable) float64 // GetPercent?
	GetMinSize(d Displayable) float64
	GetPadding(d Displayable) float64
	GetPaddingFirst(d Displayable) float64
	GetPaddingLast(d Displayable) float64
	GetPosition(d Displayable) float64
	GetPreferred(d Displayable) float64
	GetSize(d Displayable) float64
}
