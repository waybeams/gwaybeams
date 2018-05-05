package events

var lastID int64

func newHandlerID() int64 {
	lastID = lastID + 1
	return lastID
}

type Event interface {
	Name() string
	Cancel()
	IsCancelled() bool
	Payload() interface{}
	Target() interface{}
	// NOTE: Cannot support cyclic dependency, need to figure out how/where
	// to manage interfaces for this to work.
	// DisplayTarget() display.Displayable
}

type EventBase struct {
	name        string
	payload     interface{}
	target      interface{}
	isCancelled bool
}

func (e *EventBase) IsCancelled() bool {
	return e.isCancelled
}

func (e *EventBase) Cancel() {
	e.isCancelled = true
}

// func (e *EventBase) DisplayTarget() display.Displayable {
// return e.target.(display.Displayable)
// }

func (e *EventBase) Name() string {
	return e.name
}

func (e *EventBase) Payload() interface{} {
	return e.payload
}

func (e *EventBase) Target() interface{} {
	return e.target
}

type EventHandler func(e Event)

// Unsubscriber is a scoped handler removal function that will return true if the
// function was successfully removed and false if it was not found.
type Unsubscriber func() bool

type registeredHandler struct {
	eventName string
	handler   EventHandler
	id        int64
}

type Emitter interface {
	On(eventName string, handler EventHandler) Unsubscriber
	Bubble(event Event)
	Emit(event Event)
	RemoveAllHandlers() bool
	RemoveAllHandlersFor(eventName string) bool
}

type EmitterBase struct {
	handlers []*registeredHandler
}

func (e *EmitterBase) RemoveAllHandlersFor(eventName string) bool {
	var found = false
	var remaining []*registeredHandler
	for _, entry := range e.handlers {
		if entry.eventName != eventName {
			remaining = append(remaining, entry)
			found = true
		}
	}
	e.handlers = remaining
	return found
}

func (e *EmitterBase) RemoveAllHandlers() bool {
	found := len(e.handlers) > 0
	e.handlers = nil
	return found
}

func (e *EmitterBase) Bubble(event Event) {
	// NOTE(lbayes): Component overrides this method and implements support
	// that requires access to the Composable interface.
	panic("Template method should be overridden")
}

func (e *EmitterBase) On(eventName string, handler EventHandler) Unsubscriber {
	id := newHandlerID()
	rHandler := &registeredHandler{
		id:        id,
		eventName: eventName,
		handler:   handler,
	}

	e.handlers = append(e.handlers, rHandler)
	return func() bool {
		for index, entry := range e.handlers {
			if entry.id == id {
				e.handlers = append(e.handlers[:index], e.handlers[index+1:]...)
			}
			return true
		}
		return false
	}
}

func (e *EmitterBase) Emit(event Event) {
	for _, entry := range e.handlers {
		if event.IsCancelled() {
			return
		}
		if entry.eventName == event.Name() {
			entry.handler(event)
		}
	}
}

// NewEmitter creates and returns a base emitter that can be composed by any entity.
func NewEmitter() *EmitterBase {
	return &EmitterBase{}
}

// New creates an Event instance with the provided configuration.
func New(eventName string, target interface{}, payload interface{}) *EventBase {
	return &EventBase{name: eventName, target: target, payload: payload}
}
