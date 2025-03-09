package state

type Manager[T Model[T]] struct {
	state     *State[T]
	accessors map[string]*Accessor[T]
}

func NewManager[T Model[T]](state *State[T]) *Manager[T] {
	return &Manager[T]{
		state:     state,
		accessors: make(map[string]*Accessor[T]),
	}
}

func (m *Manager[T]) AccessorFor(itemKey string) (*Accessor[T], error) {
	// Check if an accessor for the itemKey already exists
	if accessor, ok := m.accessors[itemKey]; ok {
		return accessor, nil
	}

	// Otherwise, create a new accessor
	item, err := m.state.Get(itemKey)
	if err != nil {
		return nil, err
	}
	accessor := NewAccessor(item)
	m.accessors[itemKey] = accessor
	return accessor, nil
}
