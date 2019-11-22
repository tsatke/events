package events_test

import (
	"sync/atomic"
	"testing"

	"github.com/TimSatke/events"
)

func TestDispatcherPayload(t *testing.T) {
	evtType := uint16(4)

	dispatcher := events.NewDispatcher()

	source := "source"
	data := "data"

	f := func(evt events.Event) {
		if evt.Type != evtType {
			t.Fatalf("Mismatch event type (actual: %v, expected: %v)", evt.Type, evtType)
		}
		if evt.Source != source {
			t.Fatalf("Mismatch event source (actual: %v, expected: %v)", evt.Source, source)
		}
		if evt.Data != data {
			t.Fatalf("Mismatch event data (actual: %v, expected: %v)", evt.Data, data)
		}
	}

	dispatcher.RegisterFunc(evtType, f)

	evt := events.Event{
		Type:   evtType,
		Source: source,
		Data:   data,
	}

	dispatcher.Fire(evt)
}

func TestDispatcherUnregisterNil(t *testing.T) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()

	execs := uint64(0)
	f := func(_ events.Event) {
		atomic.AddUint64(&execs, 1)
	}
	consumer := events.NewConsumer(f)

	dispatcher.Register(evtType, consumer)

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	dispatcher.Fire(evt)
	dispatcher.Unregister(nil)
	dispatcher.Fire(evt)

	if execs != 2 {
		t.Fatalf("It seems like a Consumer has been unregistered (%v calls recorded)", execs)
	}
}

func TestDispatcherUnregisterConsumer(t *testing.T) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()

	execs := uint64(0)
	f := func(_ events.Event) {
		atomic.AddUint64(&execs, 1)
	}
	consumer := events.NewConsumer(f)

	dispatcher.Register(evtType, consumer)

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	dispatcher.Fire(evt)
	dispatcher.Unregister(consumer)
	dispatcher.Fire(evt)

	if execs != 1 {
		t.Fatalf("It seems like the Consumer has not been unregistered (%v calls recorded)", execs)
	}
}

func TestDispatcherRegisterConsumerTwice(t *testing.T) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()

	execs := uint64(0)
	f := func(_ events.Event) {
		atomic.AddUint64(&execs, 1)
	}
	consumer := events.NewConsumer(f)

	dispatcher.Register(evtType, consumer)
	dispatcher.Register(evtType, consumer)

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	dispatcher.Fire(evt)

	if execs != 1 {
		t.Fatalf("It seems like the Consumer has been registered more than once (%v calls recorded)", execs)
	}
}

func TestDispatcherRegisterNil(_ *testing.T) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()

	dispatcher.Register(evtType, nil)

	dispatcher.Fire(events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	})
}

func TestDispatcherRegisterNilFunc(_ *testing.T) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()

	dispatcher.RegisterFunc(evtType, nil)

	dispatcher.Fire(events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	})
}

func TestDispatcherFireFunc(t *testing.T) {
	evtType := uint16(0)

	totalExecs := uint64(100000)
	recordedExecs := uint64(0)

	f := func(_ events.Event) {
		atomic.AddUint64(&recordedExecs, 1)
	}

	dispatcher := events.NewDispatcher()

	dispatcher.RegisterFunc(evtType, f)

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	for i := uint64(0); i < totalExecs; i++ {
		dispatcher.Fire(evt)
	}

	if totalExecs != recordedExecs {
		t.Fatalf("Not all fired events were recorded (recorded: %v, actual: %v)", recordedExecs, totalExecs)
	}
}

func TestDispatcherFireConsumer(t *testing.T) {
	evtType := uint16(0)

	totalExecs := uint64(100000)
	recordedExecs := uint64(0)

	f := func(_ events.Event) {
		atomic.AddUint64(&recordedExecs, 1)
	}

	dispatcher := events.NewDispatcher()

	dispatcher.Register(evtType, events.NewConsumer(f))

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	for i := uint64(0); i < totalExecs; i++ {
		dispatcher.Fire(evt)
	}

	if totalExecs != recordedExecs {
		t.Fatalf("Not all fired events were recorded (recorded: %v, actual: %v)", recordedExecs, totalExecs)
	}
}

func TestDispatcherFireBoth(t *testing.T) {
	evtType := uint16(0)

	totalExecs := uint64(100000)
	recordedExecs := uint64(0)

	f := func(_ events.Event) {
		atomic.AddUint64(&recordedExecs, 1)
	}

	dispatcher := events.NewDispatcher()

	dispatcher.RegisterFunc(evtType, f)
	dispatcher.Register(evtType, events.NewConsumer(f))

	evt := events.Event{
		Type:   evtType,
		Source: nil,
		Data:   nil,
	}

	for i := uint64(0); i < totalExecs; i++ {
		dispatcher.Fire(evt)
	}

	if totalExecs != recordedExecs/2 {
		t.Fatalf("Not all fired events were recorded (recorded: %v, actual: %v)", recordedExecs, totalExecs)
	}
}
