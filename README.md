<p align="center">
    <h1 align="center">Events</h1>
    <p align="center">A small dispatcher framework for Go.</p>
    <p align="center">
        <a href="https://github.com/TimSatke/events/actions"><img src="https://github.com/TimSatke/events/workflows/Build/badge.svg"></a>
        <a href="https://github.com/TimSatke/events/actions"><img src="https://github.com/TimSatke/events/workflows/Tests/badge.svg"></a>
        <a href="https://github.com/TimSatke/events/actions"><img src="https://github.com/TimSatke/events/workflows/Static%20analysis/badge.svg"></a>
        <br>
        <a href="https://www.codacy.com/manual/tim.satke/events?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TimSatke/events&amp;utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/c9ea4d4267774e9eb5e72ae697b401cb"/></a>
		<a href="https://www.codacy.com/manual/tim.satke/events?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TimSatke/events&amp;utm_campaign=Badge_Coverage"><img src="https://api.codacy.com/project/badge/Coverage/c9ea4d4267774e9eb5e72ae697b401cb"/></a>
    </p>
</p>

---

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