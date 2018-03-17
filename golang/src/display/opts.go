package display

type Opts struct {
	// General
	Id          string
	Title       string
	Description string
	Styles      []string

	// Constraints layout
	FlexHeight int
	FlexWidth  int
	MaxHeight  int
	MaxWidth   int
	MinHeight  int
	MinWidth   int
	PrefHeight int
	PrefWidth  int

	// Fixed layout
	Width  int
	Height int
	X      int
	Y      int

	// Style
	BackgroundColor uint
	StrokeColor     uint
	StrokeSize      int
	CornerRadius    int
}

func (o *Opts) ApplyTo(d Sprite) {
	d.width = o.Width
	d.height = o.Height
	d.x = o.X
	d.y = o.Y
}
