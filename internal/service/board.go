package service

import (
	"errors"

	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

var (
	ErrOutOfBounds  = errors.New("out of bounds")
	ErrIllegalMove  = errors.New("illegal move")
	ErrCellOccupied = errors.New("cell occupied")
)

type BoardService struct {
	accessor *state.Accessor[*qs.Board]
}

func NewBoardService(accessor *state.Accessor[*qs.Board]) *BoardService {
	return &BoardService{accessor}
}

// PlaceTile adds a tile to the Board and returns a score for the move
func (b *BoardService) PlaceTile(tile *qs.Tile, x, y int) (int, error) {
	var (
		score int
		err   error
	)
	err = b.accessor.Update(func(b *qs.Board) error {
		if x < 0 || y < 0 || x >= len(b.Tiles) || y >= len(b.Tiles[x]) {
			return ErrOutOfBounds
		}
		if b.Tiles[x][y] != nil {
			return ErrCellOccupied
		}

		// Check that the tile can be placed in the given position
		// lineVertical := b.GetLine(x, y, qs.DirectionVertical)
		// lineHorizontal := b.GetLine(x, y, qs.DirectionHorizontal)

		// b.Tiles[x][y] = tile
		return nil
	})
	if err != nil {
		return 0, err
	}
	return score, nil
}
