package tui

import (
	"github.com/konapun/qwirkle/game"
	// FIXME: don't make internal
	"github.com/konapun/qwirkle/internal/state"
)

type Observer struct {
	renderer *Renderer
}

func NewObserver() *Observer {
	renderer := NewRenderer()
	return &Observer{
		renderer: renderer,
	}
}

func (o *Observer) Update(event *game.Event) error {
	switch event.Type {
	case game.EventTypePlayersUpdated:
		updatedPlayers := event.New.(*state.Players)
		return o.renderer.RenderPlayers(updatedPlayers)
	case game.EventTypeBoardUpdated:
		updatedBoard := event.New.(*state.Board)
		return o.renderer.RenderBoard(updatedBoard)
	case game.EventTypeTileBagUpdated:
		updatedTileBag := event.New.(*state.TileBag)
		return o.renderer.RenderTileBag(updatedTileBag)
	}
	return nil
}
