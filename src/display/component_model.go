package display

type ComponentModel struct {
	// Application
	FramesPerSecond int

	// General
	Id          string
	Title       string
	Description string

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
	LayoutType        LayoutTypeValue
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
}
