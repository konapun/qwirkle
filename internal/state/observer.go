package state

type Observer[T any] interface {
	Update(old, new T)
}

type RuntimeObserver[T any] struct {
	update func(new T, old T)
}

func NewRuntimeObserver[T any](update func(T, T)) *RuntimeObserver[T] {
	return &RuntimeObserver[T]{
		update: update,
	}
}

func (o *RuntimeObserver[T]) Update(new T, old T) {
	o.update(new, old)
}
