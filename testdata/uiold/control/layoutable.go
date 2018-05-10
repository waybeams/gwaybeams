package control

import (
	"fmt"
	"uiold/layout"
	"log"
	"math"
	"ui"
)

func (c *Control) ActualHeight() float64 {
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

func (c *Control) ActualWidth() float64 {
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

func (c *Control) SetLayoutType(layoutType ui.LayoutTypeValue) {
	c.Model().LayoutType = layoutType
}

func (c *Control) LayoutType() ui.LayoutTypeValue {
	return c.Model().LayoutType
}

func (c *Control) Layout() {
	c.GetLayout()(c)
	c.LayoutChildren()
}

func (c *Control) LayoutChildren() {
	for _, child := range c.Children() {
		child.Layout()
	}
}

func (c *Control) GetLayout() ui.LayoutHandler {
	// NOTE(lbayes): There's a naming conflict. Layout() is used above as a verb
	// and here as a noun.
	switch c.LayoutType() {
	case ui.StackLayoutType:
		return layout.StackLayout
	case ui.HorizontalFlowLayoutType:
		return layout.HorizontalFlowLayout
	case ui.VerticalFlowLayoutType:
		return layout.VerticalFlowLayout
	case ui.NoLayoutType:
		return layout.NoLayout
	default:
		msg := fmt.Sprintf("ERROR: Requested LayoutTypeValue (%v) is not supported:", c.LayoutType())
		log.Fatal(msg)
		return nil
	}
}

func (c *Control) SetModel(model *ui.Model) {
	c.model = model
}

func (c *Control) Model() *ui.Model {
	if c.model == nil {
		c.model = ui.NewModel()
	}
	return c.model
}

func (c *Control) SetX(x float64) {
	c.Model().X = x
}

func (c *Control) SetY(y float64) {
	c.Model().Y = y
}

func (c *Control) SetTextX(x float64) {
	c.Model().TextX = x
}

func (c *Control) SetTextY(y float64) {
	c.Model().TextY = y
}

func (c *Control) SetZ(z float64) {
	c.Model().Z = z
}

func (c *Control) TextX() float64 {
	return (c.X() + c.PaddingLeft()) - c.Model().TextX
}

func (c *Control) TextY() float64 {
	return (c.Y() + c.PaddingTop()) - c.Model().TextY
}

func (c *Control) X() float64 {
	return c.Model().X
}

func (c *Control) Y() float64 {
	return c.Model().Y
}

func (c *Control) Z() float64 {
	return c.Model().Z
}

func (c *Control) SetHAlign(value ui.Alignment) {
	c.Model().HAlign = value
}

func (c *Control) HAlign() ui.Alignment {
	return c.Model().HAlign
}

func (c *Control) VAlign() ui.Alignment {
	return c.Model().VAlign
}

func (c *Control) SetVAlign(value ui.Alignment) {
	c.Model().VAlign = value
}

func (c *Control) SetWidth(w float64) {
	model := c.Model()
	if model.Width != w {
		model.Width = -1
		c.SetActualWidth(w)
	}
}

func (c *Control) SetHeight(h float64) {
	model := c.Model()
	if model.Height != h {
		model.Height = -1
		c.SetActualHeight(h)
	}
}

func (c *Control) WidthInBounds(width float64) float64 {
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

func (c *Control) HeightInBounds(height float64) float64 {
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

func (c *Control) Width() float64 {
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

func (c *Control) Height() float64 {
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

func (c *Control) FixedWidth() float64 {
	return c.Model().Width
}

func (c *Control) FixedHeight() float64 {
	return c.Model().Height
}

func (c *Control) SetPrefWidth(value float64) {
	c.Model().PrefWidth = value
}

func (c *Control) SetPrefHeight(value float64) {
	c.Model().PrefHeight = value
}

func (c *Control) PrefWidth() float64 {
	return c.Model().PrefWidth
}

func (c *Control) PrefHeight() float64 {
	return c.Model().PrefHeight
}

func (c *Control) SetActualWidth(width float64) {
	inBounds := c.WidthInBounds(width)
	model := c.Model()
	model.ActualWidth = inBounds
	if model.Width != -1 && model.Width != width {
		model.Width = width
	}
}

func (c *Control) SetActualHeight(height float64) {
	inBounds := c.HeightInBounds(height)
	model := c.Model()
	model.ActualHeight = inBounds
	if model.Height != -1 && model.Height != height {
		model.Height = height
	}
}

func (c *Control) InferredMinWidth() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinWidth())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Control) InferredMinHeight() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinHeight())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Control) SetExcludeFromLayout(value bool) {
	c.Model().ExcludeFromLayout = value
}

func (c *Control) SetMinWidth(min float64) {
	c.Model().MinWidth = min
	// Ensure we're not already too small for the new min
	if c.ActualWidth() < min {
		c.SetActualWidth(min)
	}
}

func (c *Control) SetMinHeight(min float64) {
	c.Model().MinHeight = min
	// Ensure we're not already too small for the new min
	if c.ActualHeight() < min {
		c.SetActualHeight(min)
	}
}

func (c *Control) MinWidth() float64 {
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

func (c *Control) MinHeight() float64 {
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

func (c *Control) SetMaxWidth(max float64) {
	if c.Width() > max {
		c.SetWidth(max)
	}
	c.Model().MaxWidth = max
}

func (c *Control) SetMaxHeight(max float64) {
	if c.Height() > max {
		c.SetHeight(max)
	}
	c.Model().MaxHeight = max
}

func (c *Control) MaxWidth() float64 {
	return c.Model().MaxWidth
}

func (c *Control) MaxHeight() float64 {
	return c.Model().MaxHeight
}

func (c *Control) ExcludeFromLayout() bool {
	return c.Model().ExcludeFromLayout
}

func (c *Control) SetFlexWidth(value float64) {
	c.Model().FlexWidth = value
}

func (c *Control) SetFlexHeight(value float64) {
	c.Model().FlexHeight = value
}

func (c *Control) FlexWidth() float64 {
	return c.Model().FlexWidth
}

func (c *Control) FlexHeight() float64 {
	return c.Model().FlexHeight
}

func (c *Control) SetPadding(value float64) {
	c.Model().Padding = value
}

func (c *Control) SetPaddingBottom(value float64) {
	c.Model().PaddingBottom = value
}

func (c *Control) SetPaddingLeft(value float64) {
	c.Model().PaddingLeft = value
}

func (c *Control) SetPaddingRight(value float64) {
	c.Model().PaddingRight = value
}

func (c *Control) SetPaddingTop(value float64) {
	c.Model().PaddingTop = value
}

func (c *Control) Padding() float64 {
	return c.Model().Padding
}

func (c *Control) HorizontalPadding() float64 {
	return c.PaddingLeft() + c.PaddingRight()
}

func (c *Control) VerticalPadding() float64 {
	return c.PaddingTop() + c.PaddingBottom()
}

func (c *Control) getPaddingForSide(getter func() float64) float64 {
	model := c.Model()
	if getter() == -1.0 {
		if model.Padding > -1.0 {
			return model.Padding
		}
		return -1.0
	}
	return getter()
}

func (c *Control) PaddingLeft() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingLeft
	})
}

func (c *Control) PaddingRight() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingRight
	})
}

func (c *Control) PaddingBottom() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingBottom
	})
}

func (c *Control) PaddingTop() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingTop
	})
}

func (c *Control) YOffset() float64 {
	offset := c.Y()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.YOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Control) XOffset() float64 {
	offset := c.X()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.XOffset()
	}
	return math.Max(0.0, offset)
}
