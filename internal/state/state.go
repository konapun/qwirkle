package state

import (
	"github.com/konapun/statekit/state"
)

type State struct {
	inner *state.State
}

func NewState() *State {
	innerState := state.NewState(
		NewPlayers(),
		NewBoard(),
		NewTileBag(),
	)
	return &State{
		inner: innerState,
	}
}
