package display

type ComponentOption (func(d Displayable) error)

func Id(value string) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Id = value
		return nil
	}
}

func Title(value string) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Title = value
		return nil
	}
}

func ExcludeFromLayout(value bool) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ExcludeFromLayout = value
		return nil
	}
}

func ActualWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ActualWidth = value
		return nil
	}
}

func ActualHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ActualHeight = value
		return nil
	}
}

func Width(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Width = value
		return nil
	}
}

func Height(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Height = value
		return nil
	}
}

func Size(width, height float64) ComponentOption {
	return func(d Displayable) error {
		model := d.GetModel()
		model.Width = width
		model.Height = height
		return nil
	}
}

func MaxWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MaxWidth = value
		return nil
	}
}

func MaxHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MaxHeight = value
		return nil
	}
}

func MinWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MinWidth = value
		return nil
	}
}

func MinHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MinHeight = value
		return nil
	}
}

func PrefWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PrefWidth = value
		return nil
	}
}

func PrefHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PrefHeight = value
		return nil
	}
}

func FlexWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().FlexWidth = value
		return nil
	}
}

func FlexHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().FlexHeight = value
		return nil
	}
}

func HAlign(align Alignment) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().HAlign = align
		return nil
	}
}

func VAlign(align Alignment) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().VAlign = align
		return nil
	}
}

func X(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().X = pos
		return nil
	}
}

func Y(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Y = pos
		return nil
	}
}

func Z(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Z = pos
		return nil
	}
}

func LayoutType(layoutType LayoutTypeValue) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().LayoutType = layoutType
		return nil
	}
}

func Padding(value float64) ComponentOption {
	return func(d Displayable) error {
		model := d.GetModel()
		// Set the ComponentModel object directly
		if model.PaddingBottom == 0 {
			model.PaddingBottom = -1
		}
		if model.PaddingLeft == 0 {
			model.PaddingLeft = -1
		}
		if model.PaddingRight == 0 {
			model.PaddingRight = -1
		}
		if model.PaddingTop == 0 {
			model.PaddingTop = -1
		}
		model.Padding = value
		return nil
	}
}

func PaddingBottom(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingBottom = value
		return nil
	}
}

func PaddingLeft(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingLeft = value
		return nil
	}
}

func PaddingRight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingRight = value
		return nil
	}
}

func PaddingTop(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingTop = value
		return nil
	}
}

func AttrStyles(opts ...StyleOption) ComponentOption {
	styles := NewStyleDefinition()
	return func(d Displayable) error {
		for _, opt := range opts {
			opt(styles)
		}
		d.Styles(styles)
		return nil
	}
}

// Compose children onto the current component by providing a closure that
// either accepts zero arguments, or accepts a single argument which will
// be a function that, when called will invalidate the component instance
// for a future render.
func Children(composer interface{}) ComponentOption {
	return func(d Displayable) error {
		return d.Composer(composer)
	}
}
