package car

import (
	t "github.com/caner-emec/go-event/event/types"

	"time"
)

type Car struct {
	Name 		string
	Color 		string
	Year 		string
	Engine 		string
	eventbus 	t.IEventBus
}


type events struct {
	CarStartedEvent t.EventType
	CarStoppedEvent t.EventType
}

var EventsList = events{
	CarStartedEvent: "car-started-event",
	CarStoppedEvent: "car-stopped-event",
}

func NewCar(name string, color string, year string, engine string, evb t.IEventBus) *Car {
	
	return &Car{
		Name: name,
		Color: color,
		Year: year,
		Engine: engine,
		eventbus: evb,
	}
}

func (c *Car) Start() {
	c.eventbus.Invoke(string(EventsList.CarStartedEvent), t.Event{
		Name: EventsList.CarStartedEvent,
		Args: t.EventArgs{Name: EventsList.CarStartedEvent, Msgs: []string{"Name: ", c.Name, "", "Color: ", c.Color, "", "Year: ", c.Year, "", "Engine: ", c.Engine}},
		Time: time.Now(),
		Sender: c,
	})
}

func (c *Car) Stop() {
	c.eventbus.Invoke(string(EventsList.CarStoppedEvent), t.Event{
		Name: EventsList.CarStoppedEvent,
		Args: t.EventArgs{Name: EventsList.CarStoppedEvent, Msgs: []string{"Name: ", c.Name, "", "Color: ", c.Color, "", "Year: ", c.Year, "", "Engine: ", c.Engine}},
		Time: time.Now(),
		Sender: c,
	})
}

func (c *Car) RegisterEvent(event t.EventType, key string, f t.EventHandlerFcn) {
	c.eventbus.AddHandler(string(event), key, f)
}

func (c *Car) UnregisterEvent(event t.EventType, key string) {
	c.eventbus.RemoveHandler(string(event), key)
}