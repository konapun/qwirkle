package scene_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/scene"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestPlayerTurn_Run(t *testing.T) {
	stateManager := state.NewManager(state.NewState())
	gameService := service.NewGameService(stateManager)
	input := mockInput[PlayerAction]{}

	gameOverCalled := false
	gameOver := NewScene(SceneGameOver, func(controller *Controller) error {
		gameOverCalled = true
		return nil
	})
	playerTurn := NewPlayerTurn(gameService, &input)
	controller := NewController(playerTurn, gameOver)

	// Start player turn
	input.Value = PlaceTile
	controller.Transition(playerTurn.Key())
	require.False(t, gameOverCalled)
}
