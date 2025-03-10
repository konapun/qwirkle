package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
	"github.com/stretchr/testify/require"
)

func TestPlayerService_PlayTile(t *testing.T) {
	player := qs.Player{
		Hand: []qs.Tile{{Color: qs.ColorRed, Shape: qs.ShapeCircle}},
	}
	accessor := state.NewAccessor(&player)
	playerService := NewPlayerService(accessor)

	// Test playing a tile that doesn't exist in the player's hand
	err := playerService.PlayTile(&qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeSquare})
	require.Equal(t, ErrTileNotFound, err)

	// Test playing a tile that exists in the player's hand
	err = playerService.PlayTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)

	// Test that the tile was removed from the player's hand
	err = playerService.PlayTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Equal(t, ErrTileNotFound, err)
}

func TestPlayerService_SetActive(t *testing.T) {
	player := qs.Player{}
	accessor := state.NewAccessor(&player)
	playerService := NewPlayerService(accessor)

	// Test setting the player to active
	err := playerService.SetActive(true)
	require.Nil(t, err)
	require.True(t, player.IsActive)

	// Test setting the player to inactive
	err = playerService.SetActive(false)
	require.Nil(t, err)
	require.False(t, player.IsActive)
}

func TestPlayerService_DrawTile(t *testing.T) {
	player := qs.Player{}
	accessor := state.NewAccessor(&player)
	playerService := NewPlayerService(accessor)

	// Test drawing a tile when the hand is not full
	err := playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)

	// Test drawing a tile when the hand is full
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Equal(t, ErrHandFull, err)
}

func TestPlayerService_IncrementScore(t *testing.T) {
	player := qs.Player{}
	accessor := state.NewAccessor(&player)
	playerService := NewPlayerService(accessor)

	// Test incrementing the score by a positive amount
	err := playerService.IncrementScore(10)
	require.Nil(t, err)

	// Test incrementing the score by a negative amount
	err = playerService.IncrementScore(-10)
	require.Equal(t, ErrInvalidScore, err)
}
