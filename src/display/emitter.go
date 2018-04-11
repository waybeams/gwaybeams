package display

var lastId int64

func newHandlerId() int64 {
	lastId = lastId + 1
	return lastId
}

type Event interface {
	Payload() interface{}
}

type EventBase struct {
	payload interface{}
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
	Emit(eventName string, payload interface{})
}

type EmitterBase struct {
	handlers []*registeredHandler
}

func (e *EmitterBase) GetHandlersFor(eventName string) []EventHandler {
	return nil
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
	for _, entry := range e.handlers {
		if entry.eventName == eventName {
			entry.handler(NewEvent(payload))
		}
	}
}

func NewEmitter() *EmitterBase {
	return &EmitterBase{}
}

func NewEvent(payload interface{}) *EventBase {
	return &EventBase{payload: payload}
}
