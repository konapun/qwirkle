package state

type Diff struct {
	// Define the methods for the Diff type
}

// TODO: generic type for diff
type Model interface {
	Key() string
	Diff(other Model) Diff
}
