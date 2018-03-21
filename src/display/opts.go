package display

type HorizontalAlignment int

const (
	LeftAlign = iota
	RightAlign
)

type VerticalAlignment int

const (
	BottomAlign = iota
	TopAlign
)

type Opts struct {
	// Application
	FramesPerSecond int

	// General
	Description string
	Id          string
	StyleAttrs  []Attrs
	StyleName   string
	StyleNames  []string
	Title       string

	// layout
	Layout            Layout
	Disabled          bool
	ExcludeFromLayout bool
	FlexHeight        int
	FlexWidth         int
	HAlign            HorizontalAlignment
	Height            float64
	Hidden            bool
	MaxHeight         float64
	MaxWidth          float64
	MinHeight         float64
	MinWidth          float64
	PrefHeight        float64
	PrefWidth         float64
	VAlign            VerticalAlignment
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
