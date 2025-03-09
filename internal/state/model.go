package state

type Diff struct {
	// Define the methods for the Diff type
}

type Model[T any] interface {
	Key() string
	Diff(other T) Diff
	Clone() T
}
