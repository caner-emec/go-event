package main

import (
	"github.com/caner-emec/go-event/example"

	"github.com/caner-emec/go-event/example/car"
)


func main()  {
	var examples []example.IExample = []example.IExample{
		car.NewCarExample(),
	}
	
	for _, example := range examples {
		example.Init()
		example.Run()
	}
}