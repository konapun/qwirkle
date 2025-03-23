package state

import (
	"github.com/konapun/statekit/state"
)

type Manager struct {
	PlayersAccessor *state.Accessor[*Players]
	BoardAccessor   *state.Accessor[*Board]
	TileBagAccessor *state.Accessor[*TileBag]
}

func NewManager(qs *State) *Manager {
	playersAccessor, _ := state.AccessorFor[*Players](qs.inner, PlayersKey)
	boardAccessor, _ := state.AccessorFor[*Board](qs.inner, BoardKey)
	tileBagAccessor, _ := state.AccessorFor[*TileBag](qs.inner, TileBagKey)

	return &Manager{
		PlayersAccessor: playersAccessor,
		BoardAccessor:   boardAccessor,
		TileBagAccessor: tileBagAccessor,
	}
}
