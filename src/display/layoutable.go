package display

import (
	"log"
	"math"
)

// Layoutable is a set of methods for components that can be positions and
// scaled.
type Layoutable interface {
	ActualHeight() float64
	ActualWidth() float64
	ExcludeFromLayout() bool
	FixedHeight() float64
	FixedWidth() float64
	FlexHeight() float64
	FlexWidth() float64
	Gutter() float64
	HAlign() Alignment
	Height() float64
	HorizontalPadding() float64
	Layout()
	LayoutChildren()
	LayoutType() LayoutTypeValue
	MaxHeight() float64
	MaxWidth() float64
	MinHeight() float64
	MinWidth() float64
	Model() *ComponentModel
	Padding() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	PaddingTop() float64
	PrefHeight() float64
	PrefWidth() float64
	SetActualHeight(height float64)
	SetActualWidth(width float64)
	SetExcludeFromLayout(bool)
	SetFlexHeight(int float64)
	SetFlexWidth(int float64)
	SetGutter(value float64)
	SetHAlign(align Alignment)
	SetHeight(height float64)
	SetLayoutType(layoutType LayoutTypeValue)
	SetMaxHeight(h float64)
	SetMaxWidth(w float64)
	SetMinHeight(h float64)
	SetMinWidth(w float64)
	SetModel(model *ComponentModel)
	SetPadding(value float64)
	SetPaddingBottom(value float64)
	SetPaddingLeft(value float64)
	SetPaddingRight(value float64)
	SetPaddingTop(value float64)
	SetPrefHeight(value float64)
	SetPrefWidth(value float64)
	SetVAlign(align Alignment)
	SetWidth(width float64)
	SetX(x float64)
	SetY(y float64)
	SetZ(z float64)
	VAlign() Alignment
	VerticalPadding() float64
	Width() float64
	X() float64
	XOffset() float64
	Y() float64
	YOffset() float64
	Z() float64
}

func (c *Component) ActualHeight() float64 {
	model := c.Model()

	if model.Height > -1 {
		return model.Height
	} else if model.ActualHeight > -1 {
		return model.ActualHeight
	}
	prefHeight := c.PrefHeight()
	if prefHeight > -1 {
		return prefHeight
	}

	return c.MinHeight()
}

func (c *Component) ActualWidth() float64 {
	model := c.Model()

	if model.Width > -1 {
		return model.Width
	} else if model.ActualWidth > -1 {
		return model.ActualWidth
	}
	prefWidth := c.PrefWidth()
	if prefWidth > -1 {
		return prefWidth
	}

	return c.MinWidth()
}

func (c *Component) SetLayoutType(layoutType LayoutTypeValue) {
	c.Model().LayoutType = layoutType
}

func (c *Component) LayoutType() LayoutTypeValue {
	return c.Model().LayoutType
}

func (c *Component) Layout() {
	c.GetLayout()(c)
	c.LayoutChildren()
}

func (c *Component) LayoutChildren() {
	for _, child := range c.Children() {
		child.Layout()
	}
}

func (c *Component) GetLayout() LayoutHandler {
	// NOTE(lbayes): There's a naming conflict. Layout() is used above as a verb
	// and here as a noun.
	switch c.LayoutType() {
	case StackLayoutType:
		return StackLayout
	case HorizontalFlowLayoutType:
		return HorizontalFlowLayout
	case VerticalFlowLayoutType:
		return VerticalFlowLayout
	case NoLayoutType:
		return NoLayout
	default:
		log.Fatal("ERROR: Requested LayoutTypeValue (%v) is not supported:", c.LayoutType())
		return nil
	}
}

func (c *Component) SetModel(model *ComponentModel) {
	c.model = model
}

func (c *Component) Model() *ComponentModel {
	if c.model == nil {
		c.model = &ComponentModel{}
	}
	return c.model
}

func (c *Component) SetX(x float64) {
	c.Model().X = x
}

func (c *Component) SetY(y float64) {
	c.Model().Y = y
}

func (c *Component) SetZ(z float64) {
	c.Model().Z = z
}

func (c *Component) X() float64 {
	return c.Model().X
}

func (c *Component) Y() float64 {
	return c.Model().Y
}

func (c *Component) Z() float64 {
	return c.Model().Z
}

func (c *Component) SetHAlign(value Alignment) {
	c.Model().HAlign = value
}

func (c *Component) HAlign() Alignment {
	return c.Model().HAlign
}

func (c *Component) VAlign() Alignment {
	return c.Model().VAlign
}

func (c *Component) SetVAlign(value Alignment) {
	c.Model().VAlign = value
}

func (c *Component) SetWidth(w float64) {
	model := c.Model()
	if model.Width != w {
		model.Width = -1
		c.SetActualWidth(w)
	}
}

func (c *Component) SetHeight(h float64) {
	model := c.Model()
	if model.Height != h {
		model.Height = -1
		c.SetActualHeight(h)
	}
}

func (c *Component) WidthInBounds(width float64) float64 {
	min := c.MinWidth()
	max := c.MaxWidth()

	if min > -1 {
		width = math.Max(min, width)
	}

	if max > -1 {
		width = math.Min(max, width)
	}
	return width
}

func (c *Component) HeightInBounds(height float64) float64 {
	min := c.MinHeight()
	max := c.MaxHeight()

	if min > -1 {
		height = math.Max(min, height)
	}

	if max > -1 {
		height = math.Min(max, height)
	}
	return height
}

func (c *Component) Width() float64 {
	model := c.Model()
	if model.ActualWidth == -1 {
		prefWidth := c.PrefWidth()
		if prefWidth > -1 {
			return prefWidth
		}
		inBounds := c.WidthInBounds(model.Width)
		if inBounds > -1.0 {
			return inBounds
		}
		return 0
	}
	return model.ActualWidth
}

func (c *Component) Height() float64 {
	model := c.Model()
	if model.ActualHeight == -1 {
		prefHeight := c.PrefHeight()
		if prefHeight > -1 {
			return prefHeight
		}
		inBounds := c.HeightInBounds(model.Height)
		if inBounds > -1 {
			return inBounds
		}
		return 0
	}
	return model.ActualHeight
}

func (c *Component) FixedWidth() float64 {
	return c.Model().Width
}

func (c *Component) FixedHeight() float64 {
	return c.Model().Height
}

func (c *Component) SetPrefWidth(value float64) {
	c.Model().PrefWidth = value
}

func (c *Component) SetPrefHeight(value float64) {
	c.Model().PrefHeight = value
}

func (c *Component) PrefWidth() float64 {
	return c.Model().PrefWidth
}

func (c *Component) PrefHeight() float64 {
	return c.Model().PrefHeight
}

func (c *Component) SetActualWidth(width float64) {
	inBounds := c.WidthInBounds(width)
	model := c.Model()
	model.ActualWidth = inBounds
	if model.Width != -1 && model.Width != width {
		model.Width = width
	}
}

func (c *Component) SetActualHeight(height float64) {
	inBounds := c.HeightInBounds(height)
	model := c.Model()
	model.ActualHeight = inBounds
	if model.Height != -1 && model.Height != height {
		model.Height = height
	}
}

func (c *Component) InferredMinWidth() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinWidth())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Component) InferredMinHeight() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinHeight())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Component) SetExcludeFromLayout(value bool) {
	c.Model().ExcludeFromLayout = value
}

func (c *Component) SetMinWidth(min float64) {
	c.Model().MinWidth = min
	// Ensure we're not already too small for the new min
	if c.ActualWidth() < min {
		c.SetActualWidth(min)
	}
}

func (c *Component) SetMinHeight(min float64) {
	c.Model().MinHeight = min
	// Ensure we're not already too small for the new min
	if c.ActualHeight() < min {
		c.SetActualHeight(min)
	}
}

func (c *Component) MinWidth() float64 {
	model := c.Model()
	width := model.Width
	minWidth := model.MinWidth
	result := -1.0

	if width > -1.0 {
		result = width
	}
	if minWidth > -1.0 {
		result = minWidth
	}

	inferredMinWidth := c.InferredMinWidth()
	if inferredMinWidth > 0 {
		return math.Max(result, inferredMinWidth)
	}
	return result
}

func (c *Component) MinHeight() float64 {
	model := c.Model()
	height := model.Height
	minHeight := model.MinHeight
	result := -1.0

	if height > -1.0 {
		result = height
	}
	if minHeight > -1.0 {
		result = minHeight
	}

	inferredMinHeight := c.InferredMinHeight()
	if inferredMinHeight > 0.0 {
		return math.Max(result, inferredMinHeight)
	}
	return result
}

func (c *Component) SetMaxWidth(max float64) {
	if c.Width() > max {
		c.SetWidth(max)
	}
	c.Model().MaxWidth = max
}

func (c *Component) SetMaxHeight(max float64) {
	if c.Height() > max {
		c.SetHeight(max)
	}
	c.Model().MaxHeight = max
}

func (c *Component) MaxWidth() float64 {
	return c.Model().MaxWidth
}

func (c *Component) MaxHeight() float64 {
	return c.Model().MaxHeight
}

func (c *Component) ExcludeFromLayout() bool {
	return c.Model().ExcludeFromLayout
}

func (c *Component) SetFlexWidth(value float64) {
	c.Model().FlexWidth = value
}

func (c *Component) SetFlexHeight(value float64) {
	c.Model().FlexHeight = value
}

func (c *Component) FlexWidth() float64 {
	return c.Model().FlexWidth
}

func (c *Component) FlexHeight() float64 {
	return c.Model().FlexHeight
}

func (c *Component) SetPadding(value float64) {
	c.Model().Padding = value
}

func (c *Component) SetPaddingBottom(value float64) {
	c.Model().PaddingBottom = value
}

func (c *Component) SetPaddingLeft(value float64) {
	c.Model().PaddingLeft = value
}

func (c *Component) SetPaddingRight(value float64) {
	c.Model().PaddingRight = value
}

func (c *Component) SetPaddingTop(value float64) {
	c.Model().PaddingTop = value
}

func (c *Component) Padding() float64 {
	return c.Model().Padding
}

func (c *Component) HorizontalPadding() float64 {
	return c.PaddingLeft() + c.PaddingRight()
}

func (c *Component) VerticalPadding() float64 {
	return c.PaddingTop() + c.PaddingBottom()
}

func (c *Component) getPaddingForSide(getter func() float64) float64 {
	model := c.Model()
	if getter() == -1.0 {
		if model.Padding > -1.0 {
			return model.Padding
		}
		return -1.0
	}
	return getter()
}

func (c *Component) PaddingLeft() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingLeft
	})
}

func (c *Component) PaddingRight() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingRight
	})
}

func (c *Component) PaddingBottom() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingBottom
	})
}

func (c *Component) PaddingTop() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingTop
	})
}

func (c *Component) YOffset() float64 {
	offset := c.Y()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.YOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Component) XOffset() float64 {
	offset := c.X()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.XOffset()
	}
	return math.Max(0.0, offset)
}
