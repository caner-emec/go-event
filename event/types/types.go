package types

import (
	"time"
)

type Event struct {
	Name 	EventType
	Time 	time.Time
	Sender 	interface{}
	Args 	EventArgs
}

type EventArgs struct {
	Name 	EventType
	Msgs 	[]string
}

type EventHandler func(<-chan Event, func(interface{}, EventArgs))
type EventHandlerFcn func(interface{}, EventArgs)

type IEventBus interface {
	AddHandler(name string, key string, f EventHandlerFcn)
	RemoveHandler(name string, key string)
	Invoke(string, Event)
}

type EventType string

type IEventApplication interface {
	RegisterEvent(event EventType, key string, f EventHandlerFcn)
	UnregisterEvent(event EventType, key string)
}