package scene

import (
	"github.com/konapun/qwirkle/internal"
	"github.com/konapun/qwirkle/internal/service"
)

const (
	SceneStartGame = "startGame"
)

type StartGameAction int

const (
	AddPlayer StartGameAction = iota
	Start
)

type StartGame struct {
	gameService *service.GameService
	input       internal.Input[StartGameAction]
}

func NewStartGame(gameService *service.GameService, input internal.Input[StartGameAction]) *StartGame {
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
	case Start:
		numPlayers := gameService.GetNumberOfPlayers()
		if numPlayers == 0 {
			return ErrNoPlayers
		}
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
