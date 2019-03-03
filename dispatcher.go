package events

import (
	"sync"
)

func NewDispatcher() Dispatcher {
	return &disp{
		consumers: make(map[uint16][]Consumer),
		consEvts:  make(map[string][]uint16),
		funcs:     make(map[uint16][]func(Event)),
	}
}

type disp struct {
	consumersLock sync.Mutex
	// consumers associates event types (uint16) with Consumers.
	consumers map[uint16][]Consumer
	// consEvts associates Consumer IDs with the event types they consume
	consEvts map[string][]uint16

	funcsLock sync.Mutex
	funcs     map[uint16][]func(Event)
}

func (d *disp) Register(evtType uint16, c Consumer) {
	// abort if no consumer is given
	if c == nil {
		return
	}

	// do not register the same consumer twice, not even for different event types
	id := c.ID()
	if _, alreadyRegistered := d.consEvts[id]; alreadyRegistered {
		return
	}

	d.consumersLock.Lock()
	defer d.consumersLock.Unlock()

	d.consEvts[id] = append(d.consEvts[id], evtType)
	d.consumers[evtType] = append(d.consumers[evtType], c)
}

func (d *disp) Unregister(c Consumer) {
	// abort if no consumer is given
	if c == nil {
		return
	}

	d.consumersLock.Lock()
	defer d.consumersLock.Unlock()

	id := c.ID()
	if evtTypes, ok := d.consEvts[id]; ok {
		for _, evtType := range evtTypes {
			d.deleteConsumer(evtType, id)
		}
	}
}

// deleteConsumer assumes that this object is locked.
// Otherwise, this method is not thread-safe.
func (d *disp) deleteConsumer(evtType uint16, consumerID string) {
	for i, consumer := range d.consumers[evtType] {
		if consumerID == consumer.ID() {
			// remove the consumer from the consumers
			d.consumers[evtType][i] = d.consumers[evtType][len(d.consumers[evtType])-1]
			d.consumers[evtType] = d.consumers[evtType][:len(d.consumers[evtType])-1]

			// forget what event types the consumer listened to
			delete(d.consEvts, consumerID)
			break
		}
	}
}

func (d *disp) RegisterFunc(evtType uint16, f func(Event)) {
	d.funcsLock.Lock()
	defer d.funcsLock.Unlock()

	d.funcs[evtType] = append(d.funcs[evtType], f)
}

func (d *disp) Fire(evt Event) {
	for _, f := range d.funcs[evt.Type] {
		f(evt)
	}

	consumers := d.consumers[evt.Type]
	for _, c := range consumers {
		c.Consume(evt)
	}
}
