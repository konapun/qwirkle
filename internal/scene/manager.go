package scene

import (
	"github.com/konapun/qwirkle/internal/io"
	"github.com/konapun/qwirkle/internal/service"
)

type Manager struct {
	controller *Controller
}

type InputReaders struct {
	StartGameReader  io.Input[StartGameAction]
	PlayerTurnReader io.Input[PlayerAction]
	GameOverReader   io.Input[GameOverAction]
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
