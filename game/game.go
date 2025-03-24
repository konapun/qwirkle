package game

import (
	gameInput "github.com/konapun/qwirkle/game/input"
	"github.com/konapun/qwirkle/internal/scene"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
)

type Game struct {
	sceneManager *scene.Manager
}

func New(reader gameInput.StringReader, observer Observer) *Game {
	stateManager := state.NewManager(state.NewState())
	eventObserver := NewEventObserver(stateManager)
	eventObserver.Register(observer)

	gameService := service.NewGameService(stateManager)
	input := gameInput.NewInput(reader)
	sceneManager := scene.NewManager(gameService, scene.InputReaders{
		StartGameReader:  &input.StartGameReader,
		PlayerTurnReader: &input.PlayerTurnReader,
		GameOverReader:   &input.GameOverReader,
	})

	return &Game{sceneManager}
}

func (g *Game) Run() error {
	return g.sceneManager.Start()
}
