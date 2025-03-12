package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
	"github.com/stretchr/testify/require"
)

func TestBoardService_PlaceTile(t *testing.T) {
	board := qs.NewBoard()
	accessor := state.NewAccessor(board)
	boardService := NewBoardService(accessor)
	board.Tiles[[2]int{0, 0}] = &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}

	// Test placing a tile in an occupied cell
	score, err := boardService.PlaceTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}, 0, 0)
	require.Equal(t, 0, score)
	require.Equal(t, ErrIllegalMove, err)

	// Test completing a row
	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeCircle}, 0, 1)
	require.Nil(t, err)
	require.Equal(t, 2, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorOrange, Shape: qs.ShapeCircle}, 0, 2)
	require.Nil(t, err)
	require.Equal(t, 3, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorYellow, Shape: qs.ShapeCircle}, 0, 3)
	require.Nil(t, err)
	require.Equal(t, 4, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorPurple, Shape: qs.ShapeCircle}, 0, 4)
	require.Nil(t, err)
	require.Equal(t, 5, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorGreen, Shape: qs.ShapeCircle}, 0, 5)
	require.Nil(t, err)
	require.Equal(t, 12, score)

	// Test completing a column (qwirkle)
	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeCircle}, 1, 0)
	require.Nil(t, err)
	require.Equal(t, 2, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorOrange, Shape: qs.ShapeCircle}, 2, 0)
	require.Nil(t, err)
	require.Equal(t, 3, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorYellow, Shape: qs.ShapeCircle}, 3, 0)
	require.Nil(t, err)
	require.Equal(t, 4, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorPurple, Shape: qs.ShapeCircle}, 4, 0)
	require.Nil(t, err)
	require.Equal(t, 5, score)

	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorGreen, Shape: qs.ShapeCircle}, 5, 0)
	require.Nil(t, err)
	require.Equal(t, 12, score)

	// Test completing a row and column (double qwirkle)
	board.Tiles = map[[2]int]*qs.Tile{
		{0, 0}: {Color: qs.ColorRed, Shape: qs.ShapeCircle},
		{1, 0}: {Color: qs.ColorRed, Shape: qs.ShapeSquare},
		{2, 0}: {Color: qs.ColorRed, Shape: qs.ShapeClover},
		{3, 0}: {Color: qs.ColorRed, Shape: qs.ShapeDiamond},
		{4, 0}: {Color: qs.ColorRed, Shape: qs.ShapeFourPointStar},
		// Gap at 5, 0
		{5, 1}: {Color: qs.ColorRed, Shape: qs.ShapeCircle},
		{5, 2}: {Color: qs.ColorRed, Shape: qs.ShapeSquare},
		{5, 3}: {Color: qs.ColorRed, Shape: qs.ShapeClover},
		{5, 4}: {Color: qs.ColorRed, Shape: qs.ShapeDiamond},
		{5, 5}: {Color: qs.ColorRed, Shape: qs.ShapeFourPointStar},
	}
	score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeEightPointStar}, 5, 0)
	require.Nil(t, err)
	require.Equal(t, 24, score)
}
