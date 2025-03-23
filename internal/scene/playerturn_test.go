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

  // Add tiles to the bag
  err := gameService.FillTileBag([]*state.Tile{{Color: state.ColorRed, Shape: state.ShapeCircle}})

	// Add two players
	err = gameService.AddPlayer()
	require.NoError(t, err)
	err = gameService.AddPlayer()
	require.NoError(t, err)

	gameOverCalled := false
	gameOver := NewScene(SceneGameOver, func(controller *Controller) error {
		gameOverCalled = true
		return nil
	})
	playerTurn := NewPlayerTurn(gameService, &input)
	controller := NewController(playerTurn, gameOver)

	// Test placing tiles
	input.Value = PlayerAction{Type: PlaceTiles, Arguments: PlaceTilesArguments{TileRun: &state.Run{
		Direction: state.DirectionLeftToRight,
		X:         0,
		Y:         0,
		Tiles:     []*state.Tile{{Color: state.ColorRed, Shape: state.ShapeCircle}},
	}}}
	controller.Transition(playerTurn.Key())
	require.False(t, gameOverCalled)

	// Test swapping tiles
	// input.Value = PlayerAction{Type: SwapTiles, Arguments: SwapTilesArguments{Tiles: []*state.Tile{{Color: state.ColorRed, Shape: state.ShapeCircle}}}}
	// controller.Transition(playerTurn.Key())
	// require.False(t, gameOverCalled)

	// Test unknown action
	input.Value = PlayerAction{Type: 2}
	err = controller.Transition(playerTurn.Key())
	require.Error(t, err)
	require.False(t, gameOverCalled)

	// Test game over
	stateManager.PlayersAccessor.Update(func(players *state.Players) error {
		for _, player := range players.Players {
			player.Hand = []*state.Tile{}
		}
		return nil
	})
	stateManager.TileBagAccessor.Update(func(tileBag *state.TileBag) error {
		tileBag.Tiles = []*state.Tile{}
		return nil
	})
	controller.Transition(playerTurn.Key())
	require.True(t, gameOverCalled)
}
