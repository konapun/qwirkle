package game

import (
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

const (
	EventTypePlayersUpdated EventType = iota
	EventTypeBoardUpdated
	EventTypeTileBagUpdated
)

type EventType int

type Event struct {
	Type EventType
	Old  any
	New  any
}

type Observer interface {
	Update(event *Event) error
}

type aggregateObserver struct {
	observers []Observer
}

func (o *aggregateObserver) Notify(event *Event) {
	for _, observer := range o.observers {
		observer.Update(event)
	}
}

// EventObserver is a struct that aggregates events from multiple observers so that they can be subscribed to in one place.
type EventObserver struct {
	observer *aggregateObserver
}

func NewEventObserver(manager *qs.Manager) *EventObserver {
	observer := &aggregateObserver{}
	manager.PlayersAccessor.RegisterObserver(state.NewRuntimeObserver(func(new, old *qs.Players) {
		observer.Notify(&Event{
			Type: EventTypePlayersUpdated,
			Old:  old,
			New:  new,
		})
	}))
	manager.BoardAccessor.RegisterObserver(state.NewRuntimeObserver(func(new, old *qs.Board) {
		observer.Notify(&Event{
			Type: EventTypeBoardUpdated,
			Old:  old,
			New:  new,
		})
	}))
	manager.TileBagAccessor.RegisterObserver(state.NewRuntimeObserver(func(new, old *qs.TileBag) {
		observer.Notify(&Event{
			Type: EventTypeTileBagUpdated,
			Old:  old,
			New:  new,
		})
	}))

	return &EventObserver{observer}
}

func (e *EventObserver) Register(observer Observer) {
	e.observer.observers = append(e.observer.observers, observer)
}
