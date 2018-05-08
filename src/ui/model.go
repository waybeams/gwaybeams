package ui

type Model struct {
	// Application
	FramesPerSecond int

	// General
	ID          string
	Key         string
	Title       string
	Description string

	// Layout
	ActualHeight      float64
	ActualWidth       float64
	Builder           Builder
	Data              map[string]interface{}
	Disabled          bool
	ExcludeFromLayout bool
	FlexHeight        float64
	FlexWidth         float64
	Focused           bool
	Gutter            float64
	HAlign            Alignment
	Height            float64
	Hidden            bool
	IsFocusable       bool
	IsText            bool
	IsTextInput       bool
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
	Selected          bool
	Text              string
	TextX             float64
	TextY             float64
	TraitNames        []string
	TypeName          string
	VAlign            Alignment
	Visible           bool
	Width             float64
	X                 float64
	Y                 float64
	Z                 float64

	// Style
	BgColor     int
	FontColor   int
	FontFace    string
	FontSize    int
	StrokeColor int
	StrokeSize  int
}

// NewModel returns a Model instance with default values configured.
func NewModel() *Model {
	model := &Model{}
	model.ActualHeight = -1
	model.ActualWidth = -1
	model.FlexHeight = -1
	model.FlexWidth = -1
	model.HAlign = AlignLeft
	model.VAlign = AlignTop
	model.Height = -1
	model.Width = -1
	model.MaxHeight = -1
	model.MaxWidth = -1
	model.MinHeight = -1
	model.MinWidth = -1
	model.Padding = -1
	model.PaddingBottom = -1
	model.PaddingLeft = -1
	model.PaddingRight = -1
	model.PaddingTop = -1
	model.PrefHeight = -1
	model.PrefWidth = -1
	model.X = 0
	model.Y = 0
	model.Z = 0
	model.LayoutType = NoLayoutType
	model.FontColor = -1
	model.FontSize = -1
	model.BgColor = -1
	model.StrokeSize = -1
	model.StrokeColor = -1
	model.Text = ""
	model.TextX = -1
	model.TextY = -1
	return model
}
