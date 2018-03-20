package display

type Opts struct {
	// General
	Id          string
	Title       string
	Description string
	Styles      []string

	// Constraints layout
	FlexHeight float64
	FlexWidth  float64
	MaxHeight  float64
	MaxWidth   float64
	MinHeight  float64
	MinWidth   float64
	PrefHeight float64
	PrefWidth  float64

	// Fixed layout
	Width  float64
	Height float64
	X      float64
	Y      float64
	Z      float64

	// Style
	BackgroundColor uint
	CornerRadius    float64
	Disabled        bool
	Margins         float64
	LineHeight      float64
	Padding         float64
	StrokeColor     uint
	StrokeSize      float64
	Visible         bool
}
