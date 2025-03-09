package state

type State struct {
	items map[string]Model
}

func NewState(items ...Model) *State {
	mappedItems := make(map[string]Model)
	for _, item := range items {
		mappedItems[item.Key()] = item
	}
	return &State{
		items: mappedItems,
	}
}

func (s *State) Get(key string) (Model, error) {
	if item, ok := s.items[key]; ok {
		return item, nil
	}
	return nil, ErrStateItemNotFound
}
