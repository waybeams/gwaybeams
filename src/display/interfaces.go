package display

// DisplayableFilter is a function that accepts a Displayable and returns
// true if it should be included.
type DisplayableFilter = func(Displayable) bool

// Composable is a set of methods that are used for composition and tree
// traversal.
type Composable interface {
	AddChild(child Displayable) int
	Composer(composeFunc interface{}) error
	GetBuilder() Builder
	GetChildAt(index int) Displayable
	GetChildCount() int
	GetChildren() []Displayable
	GetComposeEmpty() func()
	GetComposeWithBuilder() func(Builder)
	GetComposeWithBuilderAndComponent() func(Builder, Displayable)
	GetComposeWithComponent() func(Displayable)
	GetFilteredChildren(DisplayableFilter) []Displayable
	GetID() string
	GetIsContainedBy(d Displayable) bool
	GetParent() Displayable
	GetPath() string
	GetTypeName() string
	GetXOffset() float64
	GetYOffset() float64
	TypeName(name string)

	// TODO(lbayes): This should be capitalized so that external components can implement it.
	setParent(parent Displayable)
}

// Layoutable is a set of methods for components that can be positions and
// scaled.
type Layoutable interface {
	Model(model *ComponentModel)
	GetModel() *ComponentModel
	Invalidate()

	Layout()
	LayoutChildren()

	ActualHeight(height float64)
	ActualWidth(width float64)
	ExcludeFromLayout(bool)
	FlexHeight(int float64)
	FlexWidth(int float64)
	GetActualHeight() float64
	GetActualWidth() float64
	GetExcludeFromLayout() bool
	GetFixedHeight() float64
	GetFixedWidth() float64
	GetFlexHeight() float64
	GetFlexWidth() float64
	GetHAlign() Alignment
	GetHeight() float64
	GetHorizontalPadding() float64
	GetLayoutType() LayoutTypeValue
	GetMaxWidth() float64
	GetMaxHeight() float64
	GetMinHeight() float64
	GetMinWidth() float64
	GetPadding() float64
	GetPaddingBottom() float64
	GetPaddingLeft() float64
	GetPaddingRight() float64
	GetPaddingTop() float64
	GetPrefHeight() float64
	GetPrefWidth() float64
	GetVAlign() Alignment
	GetVerticalPadding() float64
	GetWidth() float64
	GetX() float64
	GetY() float64
	GetZ() float64
	HAlign(align Alignment)
	Height(height float64)
	LayoutType(layoutType LayoutTypeValue)
	MaxHeight(h float64)
	MaxWidth(w float64)
	MinHeight(h float64)
	MinWidth(w float64)
	Padding(value float64)
	PaddingBottom(value float64)
	PaddingLeft(value float64)
	PaddingRight(value float64)
	PaddingTop(value float64)
	PrefHeight(value float64)
	PrefWidth(value float64)
	VAlign(align Alignment)
	Width(width float64)
	X(x float64)
	Y(y float64)
	Z(z float64)
}

// Styleable entities can have their visual styles updated.
type Styleable interface {
	BgColor(color int)
	FontColor(color int)
	FontFace(face string)
	FontSize(size int)
	GetBgColor() int
	GetFontColor() int
	GetFontFace() string
	GetFontSize() int
	GetStrokeColor() int
	GetStrokeSize() int
	StrokeColor(color int)
	StrokeSize(size int)
}

type Clickable interface {
	OnClick(handler EventHandler)
	Click()
}

// Displayable entities can be composed, scaled, positioned, and drawn.
type Displayable interface {
	Composable
	Layoutable
	Styleable
	Clickable

	// Text and Title are both kind of weird for the general
	// component case... Need to think more about this.
	Draw(s Surface)
	GetText() string
	GetTitle() string
	GetView() RenderHandler
	Text(text string)
	Title(title string)
	View(view RenderHandler)

	PushTrait(sel string, opts ...ComponentOption) error
	GetTraitOptions() TraitOptions
}

type Event interface {
}

// Window is an outermost component that manages the application event loop.
// Concrete Window implementations will connect the component Draw() calls with
// an appropriate native rendering surface.
type Window interface {
	Init()

	GetFrameRate() int
	GetHeight() float64
	GetTitle() string
	GetWidth() float64
	Height(h float64)
	PollEvents() []Event
	Title(str string)
	Width(w float64)
}

// Render is a function type that will draw component state onto the provided
// Surface
type RenderHandler func(s Surface, d Displayable) error

// EventHandler is a function that will be called from an event.
type EventHandler func(d Displayable)
