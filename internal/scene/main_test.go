package scene_test

type mockInput[T any] struct {
	Value T
}

func (m *mockInput[T]) Read() T {
	return m.Value
}
