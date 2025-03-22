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

// PlaceTile adds a tiles to the Board and returns a score for the move
// - Find out which starting point (if any - left, right for horizontal, top, bottom for vertical) is valid
// - Place the tiles, starting from the valid starting point and moving in the direction of the run
// - Calculate the score:
//   - For any lines created by the run NOT in the direction of the run, add to cumulative score
//   - Add the score for the run itself to the cumulative score
//   - Return the cumulative score
func (b *BoardService) PlaceTiles(run *qs.Run) (int, error) {
	var (
		score int
		err   error
	)
	err = b.accessor.Update(func(board *qs.Board) error {
		var (
			startCoord int
		)
		switch run.Direction {
		case qs.DirectionLeftToRight:
			startCoord = run.X
		case qs.DirectionRightToLeft:
			startCoord = run.X - len(run.Tiles) + 1
		case qs.DirectionDownToUp:
			startCoord = run.Y
		case qs.DirectionUpToDown:
			startCoord = run.Y - len(run.Tiles) + 1
		}
		x := run.X
		y := run.Y
		for i, tile := range run.Tiles {
			if run.Direction.Orientation() == qs.OrientationHorizontal {
				x = startCoord + i
			} else {
				y = startCoord + i
			}

			horizontalLine, verticalLine, err := board.Test(tile, x, y)
			if err != nil || !horizontalLine.IsValid() || !verticalLine.IsValid() {
				return ErrIllegalMove
			}

			// Actually place the tiles
			// TODO: make sure that if this operation is interrupted by an illegal move somewhere else that the changes are aborted
			board.Tiles[[2]int{x, y}] = tile

			// Calculate the score
			if run.Direction.Orientation() == qs.OrientationVertical || i == len(run.Tiles)-1 {
				horizontalLineLength := horizontalLine.Length()
				if horizontalLineLength == 6 { // Qwirkle!
					score += 12
				} else if horizontalLineLength > 1 {
					score += horizontalLine.Length()
				}
			}
			if run.Direction.Orientation() == qs.OrientationHorizontal || i == len(run.Tiles)-1 {
				verticalLineLength := verticalLine.Length()
				if verticalLineLength == 6 { // Qwirkle!
					score += 12
				} else if verticalLineLength > 1 {
					score += verticalLine.Length()
				}
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return score, nil
}
