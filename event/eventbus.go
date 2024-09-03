package event

import (
	"slices"

	t "github.com/caner-emec/go-event/event/types"
)

type EventBus struct {
	handlers map[string][]EventOpts
}

type EventOpts struct {
	k string
	e chan t.Event
	f t.EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventOpts),
	}
}


// AddHandler adds a new event handler to the event bus.
//
// Parameters:
// - name: the name of the event.
// - key: the key associated with the  event handler function.
// - f: the event handler function.
//
// Returns: None.
func (eb *EventBus) AddHandler(name string, key string, f t.EventHandlerFcn) {
	c := make(chan t.Event)

	eb.handlers[name] = append(eb.handlers[name], EventOpts{k: key, e: c, f: func(c <-chan t.Event, f func(interface{}, t.EventArgs)) {
		for v := range c {
			// before event handler
			f(v.Sender, v.Args)
			// after event handler
		}
	}})

	
	if idx :=slices.IndexFunc(eb.handlers[name], func(v EventOpts) bool { return v.k == key }); idx != -1 {
		go eb.handlers[name][idx].f(c, f)
	}

}

// RemoveHandler removes an event handler from the event bus.
//
// Parameters:
// - name: the name of the event.
// - key: the key associated with the event handler function.
//
// Returns: None.
func (eb *EventBus) RemoveHandler(name string, key string) {
	if idx :=slices.IndexFunc(eb.handlers[name], func(v EventOpts) bool { return v.k == key }); idx != -1 {
		close(eb.handlers[name][idx].e)
		eb.handlers[name] = slices.Delete(eb.handlers[name], idx, idx+1)
	}
}

// Invoke sends an event to all registered handlers for the given event name.
//
// Parameters:
// - name: the name of the event.
// - e: the event to be sent.
//
// Returns: None.
func (eb *EventBus) Invoke(name string, e t.Event) {
	for _, c := range eb.handlers[string(e.Name)] {
		c.e <- e
	}
}