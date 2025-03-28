package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestGameService_AddPlayer(t *testing.T) {
	stateManager := qs.NewManager(qs.NewState())
	gameService := NewGameService(stateManager)

	// Test adding a player
	err := gameService.AddPlayer()
	require.NoError(t, err)
	require.Equal(t, 1, len(stateManager.PlayersAccessor.Query().Players))
	err = gameService.AddPlayer()
	require.NoError(t, err)
	require.Equal(t, 2, len(stateManager.PlayersAccessor.Query().Players))
}

func TestGameService_GetNumberOfPlayers(t *testing.T) {
	stateManager := qs.NewManager(qs.NewState())
	gameService := NewGameService(stateManager)

	// Test adding a player
	gameService.AddPlayer()
	require.Equal(t, 1, gameService.GetNumberOfPlayers())
	gameService.AddPlayer()
	require.Equal(t, 2, gameService.GetNumberOfPlayers())
}

func TestGameService_PlaceTile(t *testing.T) {
	stateManager := qs.NewManager(qs.NewState())
	gameService := NewGameService(stateManager)
	redCircle := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	// redSquare := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeSquare}

	// Test placing a tile
	gameService.AddPlayer()
	gameService.FillTileBag([]*qs.Tile{redCircle})
	gameService.DrawTile()
	gameService.DrawTile()
	err := gameService.PlaceTiles(&qs.Run{
		Direction: qs.DirectionLeftToRight,
		X:         0,
		Y:         0,
		Tiles:     []*qs.Tile{redCircle},
	})
	require.NoError(t, err)
	activePlayer, _ := stateManager.PlayersAccessor.Query().GetActivePlayer()
	require.Equal(t, 1, activePlayer.Score)
}

func TestGameService_SwapTiles(t *testing.T) {
	stateManager := qs.NewManager(qs.NewState())
	gameService := NewGameService(stateManager)
	redCircle := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	redSquare := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeSquare}

	gameService.AddPlayer()
	gameService.FillTileBag([]*qs.Tile{redCircle, redSquare})
	require.Len(t, stateManager.TileBagAccessor.Query().Tiles, 2)
	gameService.DrawTile()
	require.Len(t, stateManager.TileBagAccessor.Query().Tiles, 1)

	activePlayer, err := stateManager.PlayersAccessor.Query().GetActivePlayer()
	require.NoError(t, err)
	require.Len(t, activePlayer.Hand, 1)

	err = gameService.SwapTiles(activePlayer.Hand)
	require.NoError(t, err)

	activePlayer, err = stateManager.PlayersAccessor.Query().GetActivePlayer()
	require.NoError(t, err)
	require.Len(t, activePlayer.Hand, 1)
	require.Len(t, stateManager.TileBagAccessor.Query().Tiles, 1)

}

func TestGameService_IsOver(t *testing.T) {
	stateManager := qs.NewManager(qs.NewState())
	gameService := NewGameService(stateManager)
	redCircle := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	redSquare := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeSquare}

	// Ensure game not over with empty bag but players with tiles
	gameService.AddPlayer()
	gameService.FillTileBag([]*qs.Tile{redCircle, redSquare})
	gameService.DrawTile()
	gameService.DrawTile()
	isOver, err := gameService.IsOver()
	require.NoError(t, err)
	require.False(t, isOver)

	// Ensure game over with empty bag and no players with Tiles
	stateManager.PlayersAccessor.Update(func(players *qs.Players) error {
		players.Players[0].Hand = []*qs.Tile{}
		return nil
	})
	isOver, err = gameService.IsOver()
	require.NoError(t, err)
	require.True(t, isOver)
}
