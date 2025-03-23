package scene_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/scene"
	"github.com/konapun/qwirkle/internal/service"
	"github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestStartGame_Run(t *testing.T) {
	stateManager := state.NewManager(state.NewState())
	gameService := service.NewGameService(stateManager)
	input := mockInput[StartGameAction]{}

	playerTurnCalled := false
	playerTurn := NewScene(ScenePlayerTurn, func(controller *Controller) error {
		playerTurnCalled = true
		return nil
	})
	startGame := NewStartGame(gameService, &input)
	controller := NewController(startGame, playerTurn)

	// Test starting with no players
	input.Value = Start
	err := controller.Transition(startGame.Key())
	require.Error(t, err)
	require.Equal(t, ErrNoPlayers, err)
	require.False(t, playerTurnCalled)

	// Add first player
	input.Value = AddPlayer
	err = controller.Transition(startGame.Key())
	require.NoError(t, err)

	// Add second player
	input.Value = AddPlayer
	err = controller.Transition(startGame.Key())
	require.NoError(t, err)

	// Start game
	input.Value = Start
	err = controller.Transition(startGame.Key())
	require.NoError(t, err)
	require.True(t, playerTurnCalled)
	require.Equal(t, 2, gameService.GetNumberOfPlayers())
	require.Equal(t, 6, len(stateManager.PlayersAccessor.Query().Players[0].Hand))
	require.Equal(t, 6, len(stateManager.PlayersAccessor.Query().Players[1].Hand))

  // Test unknown action
  input.Value = 2
  err = controller.Transition(startGame.Key())
  require.Error(t, err)
}
