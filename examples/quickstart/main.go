package main

import (
	"fmt"

	"gitlab.com/TimSatke/events"
)

// EventTypes
const (
	EventExample uint16 = iota
	EventGreeting
	EventGoodbye
)

func main() {
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
}

func HandleGreet(evt events.Event) {
	// evt.Type == EventGreeting
	fmt.Printf("Hello, %v!\n", evt.Data)
}
