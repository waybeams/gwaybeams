package display

var lastId int64

func newHandlerId() int64 {
	lastId = lastId + 1
	return lastId
}

type Event interface {
	Name() string
	Payload() interface{}
}

type EventBase struct {
	name    string
	payload interface{}
}

func (e *EventBase) Name() string {
	return e.name
}

func (e *EventBase) Payload() interface{} {
	return e.payload
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
	AddHandler(eventName string, handler EventHandler) Unsubscriber
	// AddCaptureHandler(eventName string, handler EventHandler) Unsubscriber
	Emit(eventName string, payload interface{})
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
	var found = len(e.handlers) > 0
	e.handlers = nil
	return found
}

func (e *EmitterBase) AddHandler(eventName string, handler EventHandler) Unsubscriber {
	id := newHandlerId()
	e.handlers = append(e.handlers, &registeredHandler{id: id, eventName: eventName, handler: handler})
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

func (e *EmitterBase) Emit(eventName string, payload interface{}) {
	event := NewEvent(eventName, payload)
	for _, entry := range e.handlers {
		if entry.eventName == eventName {
			entry.handler(event)
		}
	}
}

/*
func (e *EmitterBase) AddCaptureHandler(eventName string, handler EventHandler) Unsubscriber {
	id := newHandlerId()
	e.captureHandlers = append(e.handlers, &registeredHandler{id: id, eventName: eventName, handler: handler})
	return func() bool {
		for index, entry := range e.captureHandlers {
			if entry.id == id {
				e.captureHandlers = append(e.captureHandlers[:index], e.captureHandlers[index+1:]...)
			}
			return true
		}
		return false
	}
}

func (e *EmitterBase) Bubble(eventName string, payload interface{}) {
}
*/

func NewEmitter() *EmitterBase {
	return &EmitterBase{}
}

func NewEvent(eventName string, payload interface{}) *EventBase {
	return &EventBase{name: eventName, payload: payload}
}
