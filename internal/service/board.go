package service

import (
	"errors"

	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

var (
	ErrIllegalMove = errors.New("illegal move")
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
	err = b.accessor.Update(func(board *qs.Board) error {
		// Test that the tile can be placed at the given position
		horizontalLine, verticalLine, err := board.Test(tile, x, y)
		if err != nil || !horizontalLine.IsValid() || !verticalLine.IsValid() {
			return ErrIllegalMove
		}

		// Actually place the tile
		board.Tiles[[2]int{x, y}] = tile

		// Calculate the score
		horizontalLineLength := horizontalLine.Length()
		if horizontalLineLength == 6 { // Qwirkle!
			score += 12
		} else if horizontalLineLength > 1 {
			score += horizontalLine.Length()
		}
		verticalLineLength := verticalLine.Length()
		if verticalLineLength == 6 { // Qwirkle!
			score += 12
		} else if verticalLineLength > 1 {
			score += verticalLine.Length()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return score, nil
}
