package car

import (
	"github.com/caner-emec/go-event/event"
	t "github.com/caner-emec/go-event/event/types"

	"fmt"
	"time"
)

type CarExample struct {}

var (
	eventbus *event.EventBus
	car 	 *Car
)

func NewCarExample() *CarExample {
	return &CarExample{}
}

func (ce *CarExample) Init() {
	eventbus = event.NewEventBus()
	car = NewCar("BMW", "black", "2024", "Diesel", eventbus)

	car.RegisterEvent(EventsList.CarStartedEvent, "car-started-event", handleCarStartedEvent)
	car.RegisterEvent(EventsList.CarStoppedEvent, "car-stopped-event", handleCarStoppedEvent)
}

func (ce *CarExample) Run() {
	car.Start()
	car.Stop()

	time.Sleep(2 * time.Second)

	car.Start()
	car.Stop()

	time.Sleep(2 * time.Second)
}

func handleCarStartedEvent(sender interface{}, args t.EventArgs) {
	fmt.Printf("Hello from handleCarStartedEvent -- %s!\n", args)
	if car, ok := sender.(*Car); ok {
		car.Color = "blue"
	}
}

func handleCarStoppedEvent(sender interface{}, args t.EventArgs) {
	fmt.Printf("Hello from handleCarStoppedEvent -- %s!\n", args)
}

