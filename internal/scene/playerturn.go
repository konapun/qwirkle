package scene

import (
	"github.com/konapun/qwirkle/internal/io"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
)

const ScenePlayerTurn = "playerTurn"

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

type SwapTilesArguments struct {
	Tiles []*state.Tile
}

type PlayerTurn struct {
	gameService *service.GameService
	input       io.Input[PlayerAction]
}

func NewPlayerTurn(gameService *service.GameService, input io.Input[PlayerAction]) *PlayerTurn {
	return &PlayerTurn{
		gameService: gameService,
		input:       input,
	}
}

func (p *PlayerTurn) Key() string {
	return ScenePlayerTurn
}

func (p *PlayerTurn) Run(controller *Controller) error {
	// Check if the game is over
	isGameOver, err := p.gameService.IsOver()
	if err != nil {
		return err
	}
	if isGameOver {
		return controller.Transition(SceneGameOver)
	}

	// Process player action
	action := p.input.Read()
	switch action.Type {
	case PlaceTiles:
		arguments, ok := action.Arguments.(PlaceTilesArguments)
		if !ok {
			return ErrInvalidAction
		}
		if err := p.gameService.PlaceTiles(arguments.TileRun); err != nil {
			return err
		}
	case SwapTiles:
		arguments := action.Arguments.(SwapTilesArguments)
		p.gameService.SwapTiles(arguments.Tiles)
	default:
		return ErrInvalidAction
	}

	// Take player turn for next player
	if err := p.gameService.NextPlayer(); err != nil {
		return err
	}
	return controller.Transition(ScenePlayerTurn)
}
