package ui

type Focusable interface {
	Blur()
	Focus()
	NearestFocusable() Displayable
	Focused() bool
	FocusedChild() Displayable
	IsFocusable() bool
	IsText() bool
	IsTextInput() bool
	SetFocusedChild(child Displayable)
	SetIsFocusable(value bool)
	SetIsText(value bool)
	SetIsTextInput(value bool)
}
