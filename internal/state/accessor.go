package state

// Accessor is a struct that provides access to the state.
type Accessor struct {
	model     Model
	observers []Observer
}

func NewAccessor(model Model) *Accessor {
	return &Accessor{
		model:     model,
		observers: make([]Observer, 0),
	}
}

func (a *Accessor) RegisterObserver(observer Observer) {
	a.observers = append(a.observers, observer)
}

// Query returns a read-only copy of the state
func (a *Accessor) Query() Model {
	// FIXME:
	return a.model
}

func (a *Accessor) Update(modifier func(Model) error) error {
	before := a.Query()
	if err := modifier(a.model); err != nil {
		return err
	}
	after := a.Query()

	a.notifyObservers(after, before)
	return nil
}

func (a *Accessor) notifyObservers(new, old Model) {
	for _, observer := range a.observers {
		observer.Update(new, old)
	}
}
