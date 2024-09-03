# go-event

Golang package for event handling.

To use this package your class must include **"eventbus"** object ad must implement **"IEventApplication"** interface.

```go
type IEventApplication interface {
 RegisterEvent(event EventType, key string, f EventHandlerFcn)
 UnregisterEvent(event EventType, key string)
}
```

## Example: Car

**car.go**

```go
package car

import (
 t "github.com/caner-emec/go-event/event/types"

 "time"
)

type Car struct {
 Name   string
 Color   string
 Year   string
 Engine   string
 eventbus  t.IEventBus
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
```

**usage:**

```go
func main() {
 eventBus := event.NewEventBus()

 car := exp.NewCar("BMW", "black", "2020", "2", eventBus)

 car.RegisterEvent(exp.EventsList.CarStartedEvent, "car-started-event", handleCarStartedEvent)
 car.RegisterEvent(exp.EventsList.CarStoppedEvent, "car-stopped-event", handleCarStoppedEvent)

 car.Start()
 car.Stop()

 time.Sleep(2 * time.Second)
}
```

```go
func handleCarStartedEvent(sender interface{}, args t.EventArgs) {
 fmt.Printf("Hello from handleCarStartedEvent -- %s!\n", args)
 if car, ok := sender.(*exp.Car); ok {
  car.Color = "blue"
 }
}

func handleCarStoppedEvent(sender interface{}, args t.EventArgs) {
 fmt.Printf("Hello from handleCarStoppedEvent -- %s!\n", args)
}
```
