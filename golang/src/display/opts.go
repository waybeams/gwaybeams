package display

type Opts struct {
	// General
	Id          string
	Title       string
	Description string
	Styles      []string

	// Layout
	FlexHeight int
	FlexWidth  int
	Height     int
	MaxHeight  int
	MaxWidth   int
	MinHeight  int
	MinWidth   int
	PrefHeight int
	PrefWidth  int
	Width      int
}
