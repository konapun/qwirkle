package service

import (
	"errors"

	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

var (
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
		// Check if the position is already occupied
		if _, exists := b.Tiles[[2]int{x, y}]; exists {
			return ErrCellOccupied
		}

		// Check that the tile can be placed in the given position
		// lineVertical := b.GetLine(x, y, qs.DirectionVertical)
		// lineHorizontal := b.GetLine(x, y, qs.DirectionHorizontal)

		// TODO: Implement game logic for placing a tile

		// Place the tile
		b.Tiles[[2]int{x, y}] = tile
		return nil
	})
	if err != nil {
		return 0, err
	}
	return score, nil
}
