package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestBoard_GetLine(t *testing.T) {
	redCircle := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redSquare := &Tile{Color: ColorRed, Shape: ShapeSquare}
	redSquare2 := &Tile{Color: ColorRed, Shape: ShapeSquare}
	redClover := &Tile{Color: ColorRed, Shape: ShapeClover}
	redClover2 := &Tile{Color: ColorRed, Shape: ShapeClover}
	redDiamond := &Tile{Color: ColorRed, Shape: ShapeDiamond}
	redFourPointStar := &Tile{Color: ColorRed, Shape: ShapeFourPointStar}
	redEightPointStar := &Tile{Color: ColorRed, Shape: ShapeEightPointStar}

	board := NewBoard()
	board.Tiles[[2]int{0, 0}] = redCircle
	board.Tiles[[2]int{1, 0}] = redSquare
	board.Tiles[[2]int{2, 0}] = redClover
	board.Tiles[[2]int{3, 0}] = redDiamond
	board.Tiles[[2]int{4, 0}] = redFourPointStar
	board.Tiles[[2]int{5, 0}] = redEightPointStar
	board.Tiles[[2]int{0, 1}] = redSquare2
	board.Tiles[[2]int{0, 2}] = redClover2

	// Test getting a horizontal line
	line := board.GetLine(0, 0, OrientationHorizontal)
	require.Equal(t, 6, line.Length())
	tiles := line.GetTiles()
	require.Equal(t, 6, len(tiles))
	require.Equal(t, redCircle, tiles[0])
	require.Equal(t, redSquare, tiles[1])
	require.Equal(t, redClover, tiles[2])
	require.Equal(t, redDiamond, tiles[3])
	require.Equal(t, redFourPointStar, tiles[4])
	require.Equal(t, redEightPointStar, tiles[5])

	// Test getting a vertical line
	line = board.GetLine(0, 0, OrientationVertical)
	require.Equal(t, 3, line.Length())
	tiles = line.GetTiles()
	require.Equal(t, 3, len(tiles))
	require.Equal(t, redCircle, tiles[0])
	require.Equal(t, redSquare2, tiles[1])
	require.Equal(t, redClover2, tiles[2])

	// Test getting a horizontal line starting from the middle, ensuring the entire line is returned
	line = board.GetLine(0, 0, OrientationVertical)
	require.Equal(t, 3, line.Length())
	tiles = line.GetTiles()
	require.Equal(t, 3, len(tiles))
	require.Equal(t, redCircle, tiles[0])
	require.Equal(t, redSquare2, tiles[1])
	require.Equal(t, redClover2, tiles[2])

	tiles = board.GetLine(2, 0, OrientationHorizontal).GetTiles()
	require.Len(t, tiles, 6)
	require.Equal(t, redCircle, tiles[0])
	require.Equal(t, redSquare, tiles[1])
	require.Equal(t, redClover, tiles[2])
	require.Equal(t, redDiamond, tiles[3])
	require.Equal(t, redFourPointStar, tiles[4])
	require.Equal(t, redEightPointStar, tiles[5])

	// Test getting a vertical line starting from the middle, ensuring the entire line is returned
	tiles = board.GetLine(0, 1, OrientationVertical).GetTiles()
	require.Len(t, tiles, 3)
	require.Equal(t, redCircle, tiles[0])
	require.Equal(t, redSquare2, tiles[1])
	require.Equal(t, redClover2, tiles[2])
}

func TestBoard_Test(t *testing.T) {
	redCircle := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redSquare := &Tile{Color: ColorRed, Shape: ShapeSquare}
	redClover := &Tile{Color: ColorRed, Shape: ShapeClover}
	blueClover := &Tile{Color: ColorBlue, Shape: ShapeClover}
	purpleClover := &Tile{Color: ColorPurple, Shape: ShapeClover}
	redDiamond := &Tile{Color: ColorRed, Shape: ShapeDiamond}
	// blueDiamond := &Tile{Color: ColorBlue, Shape: ShapeDiamond}
	// purpleDiamond := &Tile{Color: ColorPurple, Shape: ShapeDiamond}
	redFourPointStar := &Tile{Color: ColorRed, Shape: ShapeFourPointStar}
	redEightPointStar := &Tile{Color: ColorRed, Shape: ShapeEightPointStar}

	board := NewBoard()
	board.Tiles[[2]int{0, 0}] = redCircle
	board.Tiles[[2]int{1, 0}] = redSquare
	board.Tiles[[2]int{2, 0}] = redClover
	// Gap at 2, 1
	board.Tiles[[2]int{2, 2}] = blueClover
	// Gap at 3, 0
	// Gap at 3, 1
	board.Tiles[[2]int{3, 2}] = purpleClover
	board.Tiles[[2]int{4, 0}] = redFourPointStar
	// Gap at 5, 0
	board.Tiles[[2]int{5, 1}] = redFourPointStar
	board.Tiles[[2]int{5, 2}] = redSquare

	// Test placing a tile in an occupied position
	_, _, err := board.Test(redCircle, 0, 0)
	require.Equal(t, ErrCellOccupied, err)

	// Test placing a tile in a valid position on row
	horizontal, vertical, err := board.Test(redEightPointStar, 5, 0)
	require.NoError(t, err)
	require.True(t, horizontal.IsValid())
	require.True(t, vertical.IsValid())
	require.Equal(t, 2, horizontal.Length()) // (4,0), (5,0)
	require.Equal(t, 3, vertical.Length())   // (5,0), (5,1), (5,2)

	// Test that the operation didn't modify the board
	require.Nil(t, board.Tiles[[2]int{5, 0}])

	// Test placing a tile in a valid position in a row gap
	horizontal, vertical, err = board.Test(redDiamond, 3, 0)
	require.NoError(t, err)
	require.True(t, horizontal.IsValid())
	require.True(t, vertical.IsValid())

	// Test placing a tile in a valid position on column
	horizontal, vertical, err = board.Test(redDiamond, 0, 1)
	require.NoError(t, err)
	require.True(t, horizontal.IsValid())
	require.True(t, vertical.IsValid())

	// Test placing a tile in a valid position in a column gap
	horizontal, vertical, err = board.Test(purpleClover, 2, 1)
	require.NoError(t, err)
	require.True(t, horizontal.IsValid())
	require.True(t, vertical.IsValid())

	// Test placing a tile in a valid position on a row and column gap
	horizontal, vertical, err = board.Test(redDiamond, 5, 0)
	require.NoError(t, err)
	require.True(t, horizontal.IsValid())
	require.True(t, vertical.IsValid())

	// Test placing a tile in an invalid position in a row gap

	// Test placing a tile in an invalid position in a column gap

	// Test placing a tile in an invalid position on a row and column gap, where the row is valid but column is invalid

	// Test placing a tile in an invalid position on a row and column gap, where the row is invalid but column is valid
}

func TestBoard_Key(t *testing.T) {
	board := NewBoard()
	require.Equal(t, "board", board.Key())
}

func TestBoard_Diff(t *testing.T) {
	board := NewBoard()
	board.Tiles[[2]int{0, 0}] = &Tile{Color: ColorRed, Shape: ShapeCircle}
	other := NewBoard()
	other.Tiles[[2]int{0, 0}] = &Tile{Color: ColorRed, Shape: ShapeCircle}
	other.Tiles[[2]int{1, 0}] = &Tile{Color: ColorBlue, Shape: ShapeCircle}
	diff := board.Diff(other)
	require.True(t, diff.HasChanged("Tiles"))
}

func TestBoard_Clone(t *testing.T) {
	board := NewBoard()
	board.Tiles[[2]int{0, 0}] = &Tile{Color: ColorRed, Shape: ShapeCircle}
	board.Tiles[[2]int{1, 0}] = &Tile{Color: ColorBlue, Shape: ShapeCircle}
	clone := board.Clone()
	require.Equal(t, board.Tiles, clone.(*Board).Tiles)
}

func TestLine_Length(t *testing.T) {
	line := Line([]*Cell{{}, {}})
	require.Equal(t, 2, line.Length())
}

func TestLine_Contains(t *testing.T) {
	redCircle1 := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redCircle2 := &Tile{Color: ColorRed, Shape: ShapeCircle}
	orangeCircle := &Tile{Color: ColorOrange, Shape: ShapeCircle}

	line := Line([]*Cell{{Tile: redCircle1}})
	require.True(t, line.Contains(redCircle1))
	// Should also work with value matches
	require.True(t, line.Contains(redCircle2))
	require.False(t, line.Contains(orangeCircle))
}

func TestLine_Type(t *testing.T) {
	redCircle := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redSquare := &Tile{Color: ColorRed, Shape: ShapeSquare}
	orangeCircle := &Tile{Color: ColorOrange, Shape: ShapeCircle}

	line := Line([]*Cell{{Tile: redCircle}})
	require.Equal(t, TypeUndetermined, line.Type())

	line = Line([]*Cell{{Tile: redCircle}, {Tile: redSquare}})
	require.Equal(t, TypeColorMatch, line.Type())

	line = Line([]*Cell{{Tile: redCircle}, {Tile: orangeCircle}})
	require.Equal(t, TypeShapeMatch, line.Type())
}

func TestLine_IsValid(t *testing.T) {
	redCircle := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redCircle2 := &Tile{Color: ColorRed, Shape: ShapeCircle}
	orangeCircle := &Tile{Color: ColorOrange, Shape: ShapeCircle}
	yellowCircle := &Tile{Color: ColorYellow, Shape: ShapeCircle}
	GreenCircle := &Tile{Color: ColorGreen, Shape: ShapeCircle}
	blueCircle := &Tile{Color: ColorBlue, Shape: ShapeCircle}
	purpleCircle := &Tile{Color: ColorPurple, Shape: ShapeCircle}
	redSquare := &Tile{Color: ColorRed, Shape: ShapeSquare}
	orangeSquare := &Tile{Color: ColorOrange, Shape: ShapeSquare}
	redClover := &Tile{Color: ColorRed, Shape: ShapeClover}
	redDiamond := &Tile{Color: ColorRed, Shape: ShapeDiamond}
	redFourPointStar := &Tile{Color: ColorRed, Shape: ShapeFourPointStar}
	redEightPointStar := &Tile{Color: ColorRed, Shape: ShapeEightPointStar}

	// Test a single item
	line := Line([]*Cell{{Tile: redCircle}})
	require.True(t, line.IsValid())

	// Test two items with the same color
	line = Line([]*Cell{{Tile: redCircle}, {Tile: redSquare}})
	require.True(t, line.IsValid())

	// Test two items with the same shape
	line = Line([]*Cell{{Tile: redCircle}, {Tile: orangeCircle}})
	require.True(t, line.IsValid())

	// Test two items with different colors and shapes
	line = Line([]*Cell{{Tile: redCircle}, {Tile: orangeSquare}})
	require.False(t, line.IsValid())

	// Test two items with the same color and shape
	line = Line([]*Cell{{Tile: redCircle}, {Tile: redCircle2}})
	require.False(t, line.IsValid())

	// Test a shape match line with a duplicated color
	line = Line([]*Cell{{Tile: redCircle}, {Tile: redSquare}, {Tile: redCircle}})

	// Test a color match qwirkle
	line = Line([]*Cell{{Tile: redCircle}, {Tile: redSquare}, {Tile: redClover}, {Tile: redDiamond}, {Tile: redFourPointStar}, {Tile: redEightPointStar}})
	require.True(t, line.IsValid())

	// Test a shape match qwirkle
	line = Line([]*Cell{{Tile: redCircle}, {Tile: orangeCircle}, {Tile: yellowCircle}, {Tile: GreenCircle}, {Tile: blueCircle}, {Tile: purpleCircle}})
	require.True(t, line.IsValid())
}
