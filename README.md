# Events

A small dispatcher framework for golang.

## How to use
Go get it with
```
go get github.com/TimSatke/events
```

First, define your custom events.
```go
const (
	EventExample uint16 = iota
	EventGreeting
	EventGoodbye
)
```

Second, create a dispatcher.
```go
disp := events.NewDispatcher()
disp.RegisterFunc(EventGreeting, func(evt events.Event) {
	// evt.Type == EventGreeting
	fmt.Printf("Hello, %v!\n", evt.Data)
})
```

Third, fire an event.
```go
evtGreeting := events.Event{
    Type: EventGreeting,
    Data: "World",
}
disp.Fire(evtGreeting)
```

Have a look at the [examples](https://github.com/TimSatke/events/tree/master/examples).