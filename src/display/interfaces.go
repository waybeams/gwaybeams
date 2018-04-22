package display

// CursorState is how a component responds to cursor movement.
// A cursor may be any pointing device including fingers.
type CursorState int

const (
	CursorActive = iota
	CursorHovered
	CursorPressed
	CursorDisabled
)

// Composable is a set of methods that are used for composition and tree
// traversal.
type Composable interface {
	// ID should be a tree-unique identifier and should not change
	// for a given component reference at any point in time.
	// Uniqueness constraints are not enforced at this time, but if duplicate
	// IDs are used, selectors, rendering and other features might fail in
	// unanticipated ways. We will generally default to using the first found
	// match.
	ID() string

	// Key should be unique for all children of a given parent and will be used
	// to determine whether an instance created by re-running the compose
	// function should replace or reuse an existing child.
	Key() string

	// Path returns a slash-delimited path string that is the canonical
	// location for the given component.
	Path() string

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
	IsContainedBy(d Displayable) bool
	Parent() Displayable
	QuerySelector(selector string) Displayable
	QuerySelectorAll(selector string) []Displayable
	RemoveAllChildren()
	RemoveChild(child Displayable) int
	Root() Displayable
	SetBuilder(b Builder)
	SetParent(parent Displayable)
	SetTraitNames(name ...string)
	SetTypeName(name string)
	TraitNames() []string
	TypeName() string
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
	Gutter() float64
	HAlign() Alignment
	Height() float64
	HorizontalPadding() float64
	LayoutType() LayoutTypeValue
	MaxHeight() float64
	MaxWidth() float64
	MinHeight() float64
	MinWidth() float64
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
	SetGutter(value float64)
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
	SetVisible(visible bool)
	StrokeColor() int
	StrokeSize() int
	Visible() bool
}

type Focusable interface {
	Blur()
	Focus()
	Focused() bool
	IsFocusable() bool
	Selected() bool
	SetCursorState(CursorState)
	SetIsFocusable(value bool)
	SetSelected(value bool)
}

// Displayable entities can be composed, scaled, positioned, and drawn.
type Displayable interface {
	Emitter
	Composable
	Layoutable
	Styleable
	Focusable

	// Text and Title are both kind of weird for the general
	// component case... Need to think more about this.
	Data() interface{}
	Draw(s Surface)
	InvalidNodes() []Displayable
	Invalidate()
	InvalidateChildren()
	InvalidateChildrenFor(d Displayable)
	PushTrait(sel string, opts ...ComponentOption) error
	RecomposeChildren() []Displayable
	SetData(data interface{})
	SetText(text string)
	SetTitle(title string)
	SetView(view RenderHandler)
	ShouldRecompose() bool
	Text() string
	Title() string
	TraitOptions() TraitOptions
	View() RenderHandler

	UnsubAll()
	PushUnsubscriber(Unsubscriber)
}

// Window is an outermost component that manages the application event loop.
// Concrete Window implementations will connect the component Draw() calls with
// an appropriate native rendering surface.
type Window interface {
	Displayable

	Init()
	PollEvents() []Event
}

// DisplayableFilter is a function that accepts a Displayable and returns
// true if it should be included.
type DisplayableFilter = func(Displayable) bool

// RenderHandler is a function type that will draw component state onto the provided
// Surface.
type RenderHandler func(s Surface, d Displayable) error
