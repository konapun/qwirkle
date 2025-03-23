package state

import "github.com/konapun/statekit/state"

const (
	EventTypePlayerAdded EventType = iota
)

type EventType int

type Event struct {
	Type EventType
}

type Observer interface {
	Update(event *Event)
}

type aggregateObserver struct {
	observers []Observer
}

func (o *aggregateObserver) Update(old, new state.Model) {
	// TODO
	// for _, observer := range o.observers {
	// 	observer.Update(old, new)
	// }
}

// EventAggregator is a struct that aggregates events from multiple observers so that they can be subscribed to in one place.
type EventAggregator struct {
	observer *aggregateObserver
}

func NewEventAggregator(manager *Manager) *EventAggregator {
	observer := &aggregateObserver{}
	manager.PlayersAccessor.RegisterObserver(state.NewRuntimeObserver(func(old, new *Players) {
		observer.Update(old, new)
	}))
	manager.BoardAccessor.RegisterObserver(state.NewRuntimeObserver(func(old, new *Board) {
		observer.Update(old, new)
	}))
	manager.TileBagAccessor.RegisterObserver(state.NewRuntimeObserver(func(old, new *TileBag) {
		observer.Update(old, new)
	}))

	return &EventAggregator{observer}
}

func (e *EventAggregator) Subscribe(observer state.Observer[state.Model]) {
	// TODO
	// e.observer.observers = append(e.observer.observers, observer)
}
