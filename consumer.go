package events

import (
	"github.com/google/uuid"
)

// NewConsumer creates a new ready to use consumer, that can be registered to an
// event in a dispatcher.
func NewConsumer(f func(Event)) Consumer {
	return &cons{
		id: uuid.New().String(),
		f:  f,
	}
}

type cons struct {
	id string
	f  func(Event)
}

func (c *cons) ID() string {
	return c.id
}

func (c *cons) Consume(evt Event) {
	c.f(evt)
}
