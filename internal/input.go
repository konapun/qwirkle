package internal

type Input[T any] interface {
	Read() T
}
