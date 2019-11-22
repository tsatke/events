package main

import (
	"fmt"

	"github.com/TimSatke/events"
)

// EventTypes
const (
	EventExample uint16 = iota
	EventGreeting
	EventGoodbye
)

func ExampleFire() {
	disp := events.NewDispatcher()

	disp.RegisterFunc(EventGreeting, HandleGreet)

	evtExample := events.Event{
		Type: EventExample,
		Data: "Example",
	}

	evtGreeting := events.Event{
		Type: EventGreeting,
		Data: "World",
	}

	disp.Fire(evtExample)
	disp.Fire(evtGreeting)

	//Output:
	// Hello, World!
}

func HandleGreet(evt events.Event) {
	// evt.Type == EventGreeting
	fmt.Printf("Hello, %v!\n", evt.Data)
}
