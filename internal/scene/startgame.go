package scene

import (
	"github.com/konapun/qwirkle/internal/io"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
)

const SceneStartGame = "startGame"

const (
	AddPlayer StartGameAction = iota
	Start
)

type StartGameAction int

type StartGame struct {
	gameService *service.GameService
	input       io.Input[StartGameAction]
}

func NewStartGame(gameService *service.GameService, input io.Input[StartGameAction]) *StartGame {
	return &StartGame{
		gameService: gameService,
		input:       input,
	}
}

func (s *StartGame) Key() string {
	return SceneStartGame
}

func (s *StartGame) Run(controller *Controller) error {
	gameService := s.gameService

	action := s.input.Read()
	switch action {
	case AddPlayer:
		gameService.AddPlayer()
		// Transition back to the start to add more players
		controller.Transition(SceneStartGame)
	case Start:
		numPlayers := gameService.GetNumberOfPlayers()
		if numPlayers == 0 {
			return ErrNoPlayers
		}

		// Fill the tile bag
		gameService.FillTileBag(state.AllTiles)

		// Deal tiles to players
		for range numPlayers {
			for range service.MaxHandSize {
				gameService.DrawTile()
			}
			gameService.NextPlayer()
		}
		// Start game
		controller.Transition(ScenePlayerTurn)
	default:
		return ErrInvalidAction
	}
	return nil
}
