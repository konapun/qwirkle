package game

import (
	"github.com/konapun/qwirkle/internal"
	"github.com/konapun/qwirkle/internal/scene"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
)

type Game struct {
  controller *scene.Controller
}

func New(input internal.Input[any]) *Game {
	stateManager := state.NewManager(state.NewState())
	gameService := service.NewGameService(stateManager)

  controller := scene.NewController(
    scene.NewStartGame(gameService, input.layer[LayerStartGame]),
    scene.NewPlayerTurn(gameService, input.layer[LayerPlayerTurn]),
    scene.NewGameOver(input.layer[LayerGameOver]),
  )

  return &Game{controller}
}

func (c *Game) Run() error {
  return c.controller.Transition(scene.SceneStartGame)
}
