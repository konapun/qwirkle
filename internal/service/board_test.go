package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
	"github.com/stretchr/testify/require"
)

func TestBoardService_PlaceTile(t *testing.T) {
	board := qs.Board{}
	accessor := state.NewAccessor(&board)
	boardService := NewBoardService(accessor)

	// Test placing a tile out of bounds
	score, err := boardService.PlaceTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}, -1, 0)
	require.Equal(t, 0, score)
	require.Equal(t, ErrOutOfBounds, err)

	// Test placing a tile in an occupied cell
	// score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}, 0, 0)
	// require.Equal(t, 0, score)
	// require.Equal(t, ErrCellOccupied, err)

	// Test placing a tile in a valid position
	// score, err = boardService.PlaceTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}, 0, 1)
	// require.Nil(t, err)
	// require.Equal(t, 1, score)

	// Test placing a tile in a row which already contains the tile

	// Test placing a tile of the wrong color in a color-based row

	// Test placing a tile of the wrong shape in a shape-based row

	// Test placing a tile in a column which already contains the tile

	// Test placing a tile of the wrong color in a color-based column

	// Test placing a tile of the wrong shape in a shape-based column

	// Test placing a tile in a valid position on row

	// Test placing a tile in a valid position on column

	// Test placing a tile in a valid position on a row gap

	// Test placing a tile in a valid position on a column gap

	// Test placing a tile in a valid position on a row and column gap

	// Test completing a row (qwirkle)

	// Test completing a column (qwirkle)

	// Test completing a row and column (double qwirkle)
}
