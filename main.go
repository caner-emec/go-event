package main

import (
	"github.com/caner-emec/go-event/event"
	t "github.com/caner-emec/go-event/event/types"
	exp "github.com/caner-emec/go-event/example/car"

	"fmt"
	"time"
)

// It creates a new event bus and initializes a new car with the given parameters.
// The car is then registered to handle specific events.
// The car is started and stopped twice with a 2-second delay in between.
func main() {
	eventBus := event.NewEventBus()

	car := exp.NewCar("BMW", "black", "2020", "2", eventBus)

	car.RegisterEvent(exp.EventsList.CarStartedEvent, "car-started-event", handleCarStartedEvent)
	car.RegisterEvent(exp.EventsList.CarStoppedEvent, "car-stopped-event", handleCarStoppedEvent)

	car.Start()
	car.Stop()

	time.Sleep(2 * time.Second)

	car.Start()
	car.Stop()

	time.Sleep(2 * time.Second)
}

func handleCarStartedEvent(sender interface{}, args t.EventArgs) {
	fmt.Printf("Hello from handleCarStartedEvent -- %s!\n", args)
	if car, ok := sender.(*exp.Car); ok {
		car.Color = "blue"
	}
}

func handleCarStoppedEvent(sender interface{}, args t.EventArgs) {
	fmt.Printf("Hello from handleCarStoppedEvent -- %s!\n", args)
}
