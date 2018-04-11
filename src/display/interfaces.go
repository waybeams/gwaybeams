package display

// DisplayableFilter is a function that accepts a Displayable and returns
// true if it should be included.
type DisplayableFilter = func(Displayable) bool

// Composable is a set of methods that are used for composition and tree
// traversal.
type Composable interface {
	AddChild(child Displayable) int
	Builder() Builder
	ChildAt(index int) Displayable
	ChildCount() int
	Children() []Displayable
	Composer(composeFunc interface{}) error
	FindComponentByID(id string) Displayable
	GetComposeEmpty() func()
	GetComposeWithBuilder() func(Builder)
	GetComposeWithBuilderAndComponent() func(Builder, Displayable)
	GetComposeWithComponent() func(Displayable)
	GetFilteredChildren(DisplayableFilter) []Displayable
	ID() string
	IsContainedBy(d Displayable) bool
	Parent() Displayable
	Path() string
	RemoveAllChildren()
	RemoveChild(child Displayable) int
	Root() Displayable
	SetBuilder(b Builder)
	SetParent(parent Displayable)
	SetTraitNames(name ...string)
	TraitNames() []string
	XOffset() float64
	YOffset() float64
}

// Layoutable is a set of methods for components that can be positions and
// scaled.
type Layoutable interface {
	SetModel(model *ComponentModel)
	Model() *ComponentModel

	Layout()
	LayoutChildren()

	ActualHeight() float64
	ActualWidth() float64
	ExcludeFromLayout() bool
	FixedHeight() float64
	FixedWidth() float64
	FlexHeight() float64
	FlexWidth() float64
	HAlign() Alignment
	Height() float64
	HorizontalPadding() float64
	LayoutType() LayoutTypeValue
	MaxHeight() float64
	MaxWidth() float64
	MinHeight() float64
	MinWidth() float64
	OnEnterFrame(handler func(d Displayable))
	Padding() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	PaddingTop() float64
	PrefHeight() float64
	PrefWidth() float64
	SetActualHeight(height float64)
	SetActualWidth(width float64)
	SetExcludeFromLayout(bool)
	SetFlexHeight(int float64)
	SetFlexWidth(int float64)
	SetHAlign(align Alignment)
	SetHeight(height float64)
	SetLayoutType(layoutType LayoutTypeValue)
	SetMaxHeight(h float64)
	SetMaxWidth(w float64)
	SetMinHeight(h float64)
	SetMinWidth(w float64)
	SetPadding(value float64)
	SetPaddingBottom(value float64)
	SetPaddingLeft(value float64)
	SetPaddingRight(value float64)
	SetPaddingTop(value float64)
	SetPrefHeight(value float64)
	SetPrefWidth(value float64)
	SetVAlign(align Alignment)
	SetWidth(width float64)
	SetX(x float64)
	SetY(y float64)
	SetZ(z float64)
	VAlign() Alignment
	VerticalPadding() float64
	Width() float64
	X() float64
	Y() float64
	Z() float64
}

// Styleable entities can have their visual styles updated.
type Styleable interface {
	BgColor() int
	FontColor() int
	FontFace() string
	FontSize() int
	SetBgColor(color int)
	SetFontColor(color int)
	SetFontFace(face string)
	SetFontSize(size int)
	SetStrokeColor(color int)
	SetStrokeSize(size int)
	StrokeColor() int
	StrokeSize() int
}

type Clickable interface {
	OnClick(handler DisplayEventHandler)
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
	InvalidNodes() []Displayable
	Invalidate()
	InvalidateChildren()
	InvalidateChildrenFor(d Displayable)
	PushTrait(sel string, opts ...ComponentOption) error
	RecomposeChildren() []Displayable
	SetText(text string)
	SetTitle(title string)
	SetView(view RenderHandler)
	ShouldRecompose() bool
	Text() string
	Title() string
	TraitOptions() TraitOptions
	View() RenderHandler
}

// Window is an outermost component that manages the application event loop.
// Concrete Window implementations will connect the component Draw() calls with
// an appropriate native rendering surface.
type Window interface {
	Displayable

	Init()

	FrameRate() int
	PollEvents() []Event
}

// Render is a function type that will draw component state onto the provided
// Surface
type RenderHandler func(s Surface, d Displayable) error

// DisplayEventHandler is a function that will be called from an event.
// TODO(lbayes): Remove this and replace with Event system
type DisplayEventHandler func(d Displayable)
