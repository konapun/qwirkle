package state

type Manager struct {
	state     *State
	accessors map[string]*Accessor
}

func NewManager(state *State) *Manager {
	return &Manager{
		state:     state,
		accessors: make(map[string]*Accessor),
	}
}

func (m *Manager) AccessorFor(itemKey string) (*Accessor, error) {
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
