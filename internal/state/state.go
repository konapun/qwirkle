package state

import (
	"github.com/konapun/statekit/state"
)

type State[T Model[T]] struct {
	inner *state.State[T]
}

func NewState[T Model[T]](items ...T) *State[T] {
	return &State[T]{
		inner: state.NewState(items...),
	}
}
