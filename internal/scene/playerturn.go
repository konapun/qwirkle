package scene

import (
	"github.com/konapun/qwirkle/internal"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
)

const (
	ScenePlayerTurn = "playerTurn"
)

const (
	PlaceTiles PlayerActionType = iota
	SwapTiles
)

type PlayerActionType int

type PlayerAction struct {
	Type      PlayerActionType
	Arguments any
}

type PlaceTilesArguments struct {
	TileRun *state.Run
}

type PlayerTurn struct {
	gameService *service.GameService
	input       internal.Input[PlayerAction]
}

func NewPlayerTurn(gameService *service.GameService, input internal.Input[PlayerAction]) *PlayerTurn {
	return &PlayerTurn{
		gameService: gameService,
		input:       input,
	}
}

func (p *PlayerTurn) Key() string {
	return ScenePlayerTurn
}

func (p *PlayerTurn) Run(controller *Controller) error {
	action := p.input.Read()
	switch action.Type {
	case PlaceTiles:
		arguments := action.Arguments.(PlaceTilesArguments)
		p.gameService.PlaceTiles(arguments.TileRun)
		// place tile
	case SwapTiles:
		// swap tiles
	default:
		return ErrInvalidAction
	}

	if err := p.gameService.NextPlayer(); err != nil {
		return err
	}

	// Take player turn for next player
	return controller.Transition(ScenePlayerTurn)
}
