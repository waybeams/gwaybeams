package display

import "errors"

type Opts struct {
	// Application
	FramesPerSecond int

	// General
	Description string
	Id          string
	Title       string

	// layout
	ActualHeight      float64
	ActualWidth       float64
	Disabled          bool
	ExcludeFromLayout bool
	FlexHeight        float64
	FlexWidth         float64
	HAlign            Alignment
	Height            float64
	Hidden            bool
	LayoutType        LayoutType
	MaxHeight         float64
	MaxWidth          float64
	MinHeight         float64
	MinWidth          float64
	Padding           float64
	PaddingBottom     float64
	PaddingLeft       float64
	PaddingRight      float64
	PaddingTop        float64
	PrefHeight        float64
	PrefWidth         float64
	VAlign            Alignment
	Width             float64
	X                 float64
	Y                 float64
	Z                 float64

	/*
		// Style
		BackgroundColor uint
		CornerRadius    float64
		Disabled        bool
		LineHeight      float64
		Margins         float64
		Padding         float64
		StrokeColor     uint
		StrokeSize      float64
	*/
}

// Receive a copy of an Opts object and configure initial values
// so that we can figure out when users have explicitly set values
// to zero.
func InitializeOpts(opts *Opts) (*Opts, error) {
	if opts.PaddingLeft < 0 || opts.PaddingRight < 0 ||
		opts.PaddingTop < 0 || opts.PaddingBottom < 0 ||
		opts.Padding < 0 {
		return nil, errors.New("Padding values must be equal to, or greater than zero")
	}
	// Padding is a bit complicated and interrelated. One may specify padding
	// on any of the four sides. One may also use the Padding shortcut to apply
	// padding to ALL four sides. But what should we do when a use provides both
	// values? Currently, we accept the Padding against any sides that are not
	// overridden with a more specific setting. This means that our
	// implementation must be able to tell whether a value was user-defined or
	// not. These values should ONLY ever be retrieved from a Sprite.Get___()
	// method.
	if opts.Padding > 0 {
		if opts.PaddingLeft == 0 {
			opts.PaddingLeft = -1
		}
		if opts.PaddingRight == 0 {
			opts.PaddingRight = -1
		}
		if opts.PaddingTop == 0 {
			opts.PaddingTop = -1
		}
		if opts.PaddingBottom == 0 {
			opts.PaddingBottom = -1
		}
	}
	return opts, nil
}

// Display declaration is a normalized bag of values built from the
// semantic sugar that describes the hierarchy.
type Declaration struct {
	Options            *Opts
	Data               interface{}
	Compose            func()
	ComposeWithUpdate  func(func())
	ComposeWithSurface func(s Surface)
}

// Receive the slice of arbitrary, untyped arguments from a factory function
// and convert them into a well-formed Declaration or return an error.
// Callers can provide an array of objects that include at most 3 entries.
// These entries can include zero or one Opts object, user-typed data struct,
// and zero or one of either a func() or func(func()) callback that will
// compose children on the declared Displayable.
func NewDeclaration(args []interface{}) (decl *Declaration, err error) {
	decl = &Declaration{}

	if len(args) > 3 {
		return nil, errors.New("Too many arguments sent to CreateDeclaration for component factory")
	}

	for _, entry := range args {
		switch entry.(type) {
		case *Opts:
			if decl.Options != nil {
				return nil, errors.New("Only one Opts object expected")
			}
			decl.Options, err = InitializeOpts(entry.(*Opts))
			if err != nil {
				return nil, err
			}
		case func(Surface):
			if decl.ComposeWithSurface != nil {
				return nil, errors.New("Only one ComposeWithSurface function expected")
			}
			decl.ComposeWithSurface = entry.(func(Surface))
		case func():
			if decl.Compose != nil {
				return nil, errors.New("Only one Compose function expected")
			}
			decl.Compose = entry.(func())
		case func(func()):
			if decl.ComposeWithUpdate != nil {
				return nil, errors.New("Only one ComposeWithUpdate function expected")
			}
			decl.ComposeWithUpdate = entry.(func(func()))
		default:
			if decl.Data != nil {
				return nil, errors.New("Only one bag of component data expected")
			}
			decl.Data = entry
		}
	}

	if decl.Compose != nil && decl.ComposeWithUpdate != nil {
		return nil, errors.New("Only one composition function allowed")
	}

	if decl.Options == nil {
		decl.Options = &Opts{}
	}

	return decl, nil
}
