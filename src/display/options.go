package display

type ComponentOption (func(d Displayable) error)

func FlexWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.FlexWidth(value)
		return nil
	}
}

func FlexHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.FlexHeight(value)
		return nil
	}
}

func Padding(value float64) ComponentOption {
	return func(d Displayable) error {
		opts := d.GetOptions()
		if opts.PaddingBottom == 0 {
			opts.PaddingBottom = -1
		}
		if opts.PaddingLeft == 0 {
			opts.PaddingLeft = -1
		}
		if opts.PaddingRight == 0 {
			opts.PaddingRight = -1
		}
		if opts.PaddingTop == 0 {
			opts.PaddingTop = -1
		}
		opts.Padding = value
		return nil
	}
}

func PaddingBottom(value float64) ComponentOption {
	return func(d Displayable) error {
		d.PaddingBottom(value)
		return nil
	}
}

func PaddingLeft(value float64) ComponentOption {
	return func(d Displayable) error {
		d.PaddingLeft(value)
		return nil
	}
}

func PaddingRight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.PaddingRight(value)
		return nil
	}
}

func PaddingTop(value float64) ComponentOption {
	return func(d Displayable) error {
		d.PaddingTop(value)
		return nil
	}
}
