package events

type Event struct {
	Type   uint16
	Source interface{}
	Data   interface{}
}

type Dispatcher interface {
	Register(uint16, Consumer)
	Unregister(Consumer)

	RegisterFunc(uint16, func(Event))

	Fire(Event)
}

type Consumer interface {
	ID() string
	Consume(Event)
}
