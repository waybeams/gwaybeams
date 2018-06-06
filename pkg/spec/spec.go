package spec

import (
	"github.com/waybeams/waybeams/pkg/events"
)

type Factory func() ReadWriter

type Option func(w ReadWriter)

type RenderHandler func(s Surface, r Reader)

type Reader interface {
	events.Emitter
	StyleableReader
	FocusableReader
	ComposableReader
	LayoutableReader
	StatefulReader

	Invalidate()
	Factory() func() ReadWriter
	SiblingsFactory() func() []ReadWriter
	Text() string
	View() RenderHandler
}

type Writer interface {
	StyleableWriter
	FocusableWriter
	ComposableWriter
	LayoutableWriter
	StatefulWriter

	SetFactory(func() ReadWriter)
	SetSiblingsFactory(func() []ReadWriter)
	PushUnsub(events.Unsubscriber)
	SetText(text string)
	SetView(view RenderHandler)
	UnsubAll()
}

type ReadWriter interface {
	Reader
	Writer
}

// Spec is an internallly configurable, externally read-only bag of
// state that describes a user interface element.
type Spec struct {
	events.EmitterBase

	actualHeight      float64
	actualWidth       float64
	bgColor           uint
	children          []ReadWriter
	childrenHeight    float64
	childrenWidth     float64
	composer          interface{}
	contentHeight     float64
	contentWidth      float64
	currentState      string
	excludeFromLayout bool
	factory           func() ReadWriter
	flexHeight        float64
	flexWidth         float64
	fontColor         uint
	fontFace          string
	fontSize          float64
	gutter            float64
	hAlign            Alignment
	height            float64
	isFocusable       bool
	isFocused         bool
	isInvisible       bool
	isMeasured        bool
	isText            bool
	isTextInput       bool
	key               string
	layoutType        LayoutTypeValue
	maxHeight         float64
	maxWidth          float64
	minHeight         float64
	minWidth          float64
	name              string
	paddingBottom     float64
	paddingLeft       float64
	paddingRight      float64
	paddingTop        float64
	parent            ReadWriter
	prefHeight        float64
	prefWidth         float64
	siblingsFactory   func() []ReadWriter
	specName          string
	states            map[string][]Option
	strokeColor       uint
	strokeSize        float64
	text              string
	textX             float64
	textY             float64
	unsubs            []events.Unsubscriber
	vAlign            Alignment
	view              RenderHandler
	width             float64
	x                 float64
	y                 float64
}

func (c *Spec) Invalidate() {
	// NOTE(lbayes): The reference to "c" here, will not be expected for any embedding components,
	// so sending nil target to avoid confusion.
	c.Bubble(events.New(events.Invalidated, nil, nil))
}

// Factory return the factory function that created this node. This function would
// have been sent to a Childf(fn) call on the parent node.
func (c *Spec) Factory() func() ReadWriter {
	return c.factory
}

// SiblingsFactory returns the factory function that created this node and it's
// siblings. This function would have been sent to a Childrenf(fn) call on the
// parent node.
func (c *Spec) SiblingsFactory() func() []ReadWriter {
	return c.siblingsFactory
}

func (c *Spec) Text() string {
	return c.text
}

func (c *Spec) SetFactory(factory func() ReadWriter) {
	c.factory = factory
}

func (c *Spec) SetSiblingsFactory(factory func() []ReadWriter) {
	c.siblingsFactory = factory
}

func (c *Spec) SetText(text string) {
	c.text = text
}

func (c *Spec) SetView(view RenderHandler) {
	c.view = view
}

func (c *Spec) View() RenderHandler {
	return c.view
}

func (c *Spec) PushUnsub(unsub events.Unsubscriber) {
	c.unsubs = append(c.unsubs, unsub)
}

func (c *Spec) UnsubAll() {
	for _, unsub := range c.unsubs {
		unsub()
	}
}

func (c *Spec) Bubble(event events.Event) {
	c.Emit(event)

	current := c.Parent()
	for current != nil {
		if event.IsCancelled() {
			return
		}
		current.Emit(event)
		current = current.Parent()
	}
}

// New creates a new Spec instance.
func New() *Spec {
	return &Spec{}
}
