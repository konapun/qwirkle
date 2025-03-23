package scene_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/scene"
	"github.com/stretchr/testify/require"
)

func TestGameOver_Run(t *testing.T) {
	input := mockInput[GameOverAction]{}

	startGameCalled := false
	startGame := NewScene(SceneStartGame, func(controller *Controller) error {
		startGameCalled = true
		return nil
	})

	gameOver := NewGameOver(&input)
	controller := NewController(startGame, gameOver)

	// Test starting a new game
	input.Value = NewGame
	err := controller.Transition(gameOver.Key())
	require.NoError(t, err)
	require.True(t, startGameCalled)

	// Test quitting the game
	startGameCalled = false
	input.Value = Quit
	err = controller.Transition(gameOver.Key())
	require.NoError(t, err)
	require.False(t, startGameCalled)

  // Test unknown action
  input.Value = 2
  err = controller.Transition(gameOver.Key())
  require.Error(t, err)
}
