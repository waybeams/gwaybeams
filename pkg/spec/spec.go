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

	Text() string
	View() RenderHandler
}

type Writer interface {
	StyleableWriter
	FocusableWriter
	ComposableWriter
	LayoutableWriter
	StatefulWriter

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
	composer          interface{}
	currentState      string
	excludeFromLayout bool
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
	isInvisible       bool
	width             float64
	x                 float64
	y                 float64
}

func (c *Spec) Text() string {
	return c.text
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
