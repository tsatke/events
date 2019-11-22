package events

// Event represents an event that can be fired through a dispatcher and consumed
// with a consumer. An event consists of a type, an event source (usually
// something that caused the event to be fired) and data, which is an arbitrary
// payload.
type Event struct {
	Type   uint16
	Source interface{}
	Data   interface{}
}

// Dispatcher declares methods that a dispatcher must implement.
type Dispatcher interface {
	Register(uint16, Consumer)
	Unregister(Consumer)

	RegisterFunc(uint16, func(Event))

	Fire(Event)
}

// Consumer declares methods that a consumer must implement.
type Consumer interface {
	ID() string
	Consume(Event)
}
