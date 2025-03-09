package state

type State[T Model[T]] struct {
	items map[string]T
}

func NewState[T Model[T]](items ...T) *State[T] {
	mappedItems := make(map[string]T)
	for _, item := range items {
		mappedItems[item.Key()] = item
	}
	return &State[T]{
		items: mappedItems,
	}
}

func (s *State[T]) Get(key string) (T, error) {
	if item, ok := s.items[key]; ok {
		return item, nil
	}
	var zero T
	return zero, ErrStateItemNotFound
}
