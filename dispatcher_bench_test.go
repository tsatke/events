package events_test

import (
	"testing"

	"gitlab.com/TimSatke/events"
)

func BenchmarkDispatcherFireFunc(b *testing.B) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()
	dispatcher.RegisterFunc(evtType, func(evt events.Event) {
		_ = 0
	})

	event := events.Event{
		Type: evtType,
		Data: nil,
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dispatcher.Fire(event)
		}
	})
}

func BenchmarkDispatcherFireConsumer(b *testing.B) {
	evtType := uint16(0)

	dispatcher := events.NewDispatcher()
	dispatcher.Register(evtType, events.NewConsumer(func(evt events.Event) {
		_ = 0
	}))

	event := events.Event{
		Type: evtType,
		Data: nil,
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dispatcher.Fire(event)
		}
	})
}

func BenchmarkDispatcherFireBoth(b *testing.B) {
	evtType := uint16(0)

	f := func(evt events.Event) {
		_ = 0
	}

	dispatcher := events.NewDispatcher()
	dispatcher.Register(evtType, events.NewConsumer(f))
	dispatcher.RegisterFunc(evtType, f)

	event := events.Event{
		Type: evtType,
		Data: nil,
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dispatcher.Fire(event)
		}
	})
}
