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
	cons := &MyConsumer{}

	disp.Register(EventGreeting, cons)

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

	disp.Unregister(cons)

	disp.Fire(evtExample)
	disp.Fire(evtGreeting)
}

type MyConsumer struct {
}

func (c *MyConsumer) ID() string {
	return "cons001a"
}

func (c *MyConsumer) Consume(evt events.Event) {
	// evt.Type == EventGreeting
	fmt.Printf("Hello, %v!\n", evt.Data)
}
