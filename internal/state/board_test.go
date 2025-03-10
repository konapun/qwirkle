package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestBoard_GetLine(t *testing.T) {
	board := Board{
		Tiles: [][]*Tile{
			{{Color: ColorRed, Shape: ShapeCircle}},
			{{Color: ColorRed, Shape: ShapeSquare}},
			{{Color: ColorRed, Shape: ShapeClover}},
			{{Color: ColorRed, Shape: ShapeDiamond}},
			{{Color: ColorRed, Shape: ShapeFourPointStar}},
			{{Color: ColorRed, Shape: ShapeEightPointStar}},
		},
	}

	line := board.GetLine(0, 0, DirectionHorizontal)
	require.Equal(t, 2, line.Length())
}

func TestBoard_Key(t *testing.T) {
	board := Board{}
	require.Equal(t, "board", board.Key())
}

func TestBoard_Diff(t *testing.T) {
	board := Board{Tiles: [][]*Tile{{}, {}}}
	other := Board{Tiles: [][]*Tile{{}, {}, {}}}
	diff := board.Diff(&other)
	require.True(t, diff.HasChanged("Tiles"))
}

func TestBoard_Clone(t *testing.T) {
	board := Board{Tiles: [][]*Tile{
		{{Color: ColorRed, Shape: ShapeCircle}},
		{{Color: ColorBlue, Shape: ShapeCircle}},
	}}
	clone := board.Clone()
	require.Equal(t, board.Tiles, clone.Tiles)
}

func TestLine_Length(t *testing.T) {
	line := Line([]*Tile{{}, {}})
	require.Equal(t, 2, line.Length())
}

func TestLine_Contains(t *testing.T) {
	redCircle1 := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redCircle2 := &Tile{Color: ColorRed, Shape: ShapeCircle}
	orangeCircle := &Tile{Color: ColorOrange, Shape: ShapeCircle}

	line := Line([]*Tile{redCircle1})
	require.True(t, line.Contains(redCircle1))
	// Should also work with value matches
	require.True(t, line.Contains(redCircle2))
	require.False(t, line.Contains(orangeCircle))
}

func TestLine_Type(t *testing.T) {
	redCircle := &Tile{Color: ColorRed, Shape: ShapeCircle}
	redSquare := &Tile{Color: ColorRed, Shape: ShapeSquare}
	orangeCircle := &Tile{Color: ColorOrange, Shape: ShapeCircle}

	line := Line([]*Tile{redCircle})
	require.Equal(t, TypeUndetermined, line.Type())

	line = Line([]*Tile{redCircle, redSquare})
	require.Equal(t, TypeColorMatch, line.Type())

	line = Line([]*Tile{redCircle, orangeCircle})
	require.Equal(t, TypeShapeMatch, line.Type())
}
