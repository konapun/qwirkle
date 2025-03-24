package scene

import (
	"github.com/konapun/qwirkle/internal"
	"github.com/konapun/qwirkle/internal/service"
)

type Manager struct {
	controller *Controller
}

type InputReaders struct {
	StartGameReader  internal.Input[StartGameAction]
	PlayerTurnReader internal.Input[PlayerAction]
	GameOverReader   internal.Input[GameOverAction]
}

func NewManager(gameService *service.GameService, readers InputReaders) *Manager {
	controller := NewController(
		NewStartGame(gameService, readers.StartGameReader),
		NewPlayerTurn(gameService, readers.PlayerTurnReader),
		NewGameOver(readers.GameOverReader),
	)
	return &Manager{controller}
}

func (m *Manager) Start() error {
	return m.controller.Transition(SceneStartGame)
}
