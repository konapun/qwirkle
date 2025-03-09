package state

// Accessor is a struct that provides access to the state.
type Accessor[T Model[T]] struct {
	model     T
	observers []Observer[T]
}

func NewAccessor[T Model[T]](model T) *Accessor[T] {
	return &Accessor[T]{
		model:     model,
		observers: make([]Observer[T], 0),
	}
}

func (a *Accessor[T]) RegisterObserver(observer Observer[T]) {
	a.observers = append(a.observers, observer)
}

func (a *Accessor[T]) Update(modifier func(T) error) error {
	before := a.model.Clone()
	if err := modifier(a.model); err != nil {
		return err
	}
	after := a.model.Clone()

	a.notifyObservers(after, before)
	return nil
}

func (a *Accessor[T]) notifyObservers(new, old T) {
	for _, observer := range a.observers {
		observer.Update(new, old)
	}
}
