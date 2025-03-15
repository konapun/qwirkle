package state

import (
	"github.com/konapun/statekit/state"
)

const (
	TileBagKey = "tileBag"
)

const (
	ShapeCircle Shape = iota
	ShapeSquare
	ShapeClover
	ShapeDiamond
	ShapeFourPointStar
	ShapeEightPointStar
)

const (
	ColorRed Color = iota
	ColorOrange
	ColorYellow
	ColorGreen
	ColorBlue
	ColorPurple
)

var AllTiles = []*Tile{
	// Circles
	{Shape: ShapeCircle, Color: ColorRed},
	{Shape: ShapeCircle, Color: ColorRed},
	{Shape: ShapeCircle, Color: ColorRed},

	{Shape: ShapeCircle, Color: ColorOrange},
	{Shape: ShapeCircle, Color: ColorOrange},
	{Shape: ShapeCircle, Color: ColorOrange},

	{Shape: ShapeCircle, Color: ColorYellow},
	{Shape: ShapeCircle, Color: ColorYellow},
	{Shape: ShapeCircle, Color: ColorYellow},

	{Shape: ShapeCircle, Color: ColorGreen},
	{Shape: ShapeCircle, Color: ColorGreen},
	{Shape: ShapeCircle, Color: ColorGreen},

	{Shape: ShapeCircle, Color: ColorBlue},
	{Shape: ShapeCircle, Color: ColorBlue},
	{Shape: ShapeCircle, Color: ColorBlue},

	{Shape: ShapeCircle, Color: ColorPurple},
	{Shape: ShapeCircle, Color: ColorPurple},
	{Shape: ShapeCircle, Color: ColorPurple},

	// Squares
	{Shape: ShapeSquare, Color: ColorRed},
	{Shape: ShapeSquare, Color: ColorRed},
	{Shape: ShapeSquare, Color: ColorRed},

	{Shape: ShapeSquare, Color: ColorOrange},
	{Shape: ShapeSquare, Color: ColorOrange},
	{Shape: ShapeSquare, Color: ColorOrange},

	{Shape: ShapeSquare, Color: ColorYellow},
	{Shape: ShapeSquare, Color: ColorYellow},
	{Shape: ShapeSquare, Color: ColorYellow},

	{Shape: ShapeSquare, Color: ColorGreen},
	{Shape: ShapeSquare, Color: ColorGreen},
	{Shape: ShapeSquare, Color: ColorGreen},

	{Shape: ShapeSquare, Color: ColorBlue},
	{Shape: ShapeSquare, Color: ColorBlue},
	{Shape: ShapeSquare, Color: ColorBlue},

	{Shape: ShapeSquare, Color: ColorPurple},
	{Shape: ShapeSquare, Color: ColorPurple},
	{Shape: ShapeSquare, Color: ColorPurple},

	// Clovers
	{Shape: ShapeClover, Color: ColorRed},
	{Shape: ShapeClover, Color: ColorRed},
	{Shape: ShapeClover, Color: ColorRed},

	{Shape: ShapeClover, Color: ColorOrange},
	{Shape: ShapeClover, Color: ColorOrange},
	{Shape: ShapeClover, Color: ColorOrange},

	{Shape: ShapeClover, Color: ColorYellow},
	{Shape: ShapeClover, Color: ColorYellow},
	{Shape: ShapeClover, Color: ColorYellow},

	{Shape: ShapeClover, Color: ColorGreen},
	{Shape: ShapeClover, Color: ColorGreen},
	{Shape: ShapeClover, Color: ColorGreen},

	{Shape: ShapeClover, Color: ColorBlue},
	{Shape: ShapeClover, Color: ColorBlue},
	{Shape: ShapeClover, Color: ColorBlue},

	{Shape: ShapeClover, Color: ColorPurple},
	{Shape: ShapeClover, Color: ColorPurple},
	{Shape: ShapeClover, Color: ColorPurple},

	// Diamonds
	{Shape: ShapeDiamond, Color: ColorRed},
	{Shape: ShapeDiamond, Color: ColorRed},
	{Shape: ShapeDiamond, Color: ColorRed},

	{Shape: ShapeDiamond, Color: ColorOrange},
	{Shape: ShapeDiamond, Color: ColorOrange},
	{Shape: ShapeDiamond, Color: ColorOrange},

	{Shape: ShapeDiamond, Color: ColorYellow},
	{Shape: ShapeDiamond, Color: ColorYellow},
	{Shape: ShapeDiamond, Color: ColorYellow},

	{Shape: ShapeDiamond, Color: ColorGreen},
	{Shape: ShapeDiamond, Color: ColorGreen},
	{Shape: ShapeDiamond, Color: ColorGreen},

	{Shape: ShapeDiamond, Color: ColorBlue},
	{Shape: ShapeDiamond, Color: ColorBlue},
	{Shape: ShapeDiamond, Color: ColorBlue},

	{Shape: ShapeDiamond, Color: ColorPurple},
	{Shape: ShapeDiamond, Color: ColorPurple},
	{Shape: ShapeDiamond, Color: ColorPurple},

	// Four Point Stars
	{Shape: ShapeFourPointStar, Color: ColorRed},
	{Shape: ShapeFourPointStar, Color: ColorRed},
	{Shape: ShapeFourPointStar, Color: ColorRed},

	{Shape: ShapeFourPointStar, Color: ColorOrange},
	{Shape: ShapeFourPointStar, Color: ColorOrange},
	{Shape: ShapeFourPointStar, Color: ColorOrange},

	{Shape: ShapeFourPointStar, Color: ColorYellow},
	{Shape: ShapeFourPointStar, Color: ColorYellow},
	{Shape: ShapeFourPointStar, Color: ColorYellow},

	{Shape: ShapeFourPointStar, Color: ColorGreen},
	{Shape: ShapeFourPointStar, Color: ColorGreen},
	{Shape: ShapeFourPointStar, Color: ColorGreen},

	{Shape: ShapeFourPointStar, Color: ColorBlue},
	{Shape: ShapeFourPointStar, Color: ColorBlue},
	{Shape: ShapeFourPointStar, Color: ColorBlue},

	{Shape: ShapeFourPointStar, Color: ColorPurple},
	{Shape: ShapeFourPointStar, Color: ColorPurple},
	{Shape: ShapeFourPointStar, Color: ColorPurple},

	// Eight Point Stars
	{Shape: ShapeEightPointStar, Color: ColorRed},
	{Shape: ShapeEightPointStar, Color: ColorRed},
	{Shape: ShapeEightPointStar, Color: ColorRed},

	{Shape: ShapeEightPointStar, Color: ColorOrange},
	{Shape: ShapeEightPointStar, Color: ColorOrange},
	{Shape: ShapeEightPointStar, Color: ColorOrange},

	{Shape: ShapeEightPointStar, Color: ColorYellow},
	{Shape: ShapeEightPointStar, Color: ColorYellow},
	{Shape: ShapeEightPointStar, Color: ColorYellow},

	{Shape: ShapeEightPointStar, Color: ColorGreen},
	{Shape: ShapeEightPointStar, Color: ColorGreen},
	{Shape: ShapeEightPointStar, Color: ColorGreen},

	{Shape: ShapeEightPointStar, Color: ColorBlue},
	{Shape: ShapeEightPointStar, Color: ColorBlue},
	{Shape: ShapeEightPointStar, Color: ColorBlue},

	{Shape: ShapeEightPointStar, Color: ColorPurple},
	{Shape: ShapeEightPointStar, Color: ColorPurple},
	{Shape: ShapeEightPointStar, Color: ColorPurple},
}

type Shape int

type Color int

type Tile struct {
	Shape Shape
	Color Color
}

func (t *Tile) Equals(other *Tile) bool {
	return t.Shape == other.Shape && t.Color == other.Color
}

type TileBag struct {
	Tiles []*Tile
}

func NewTileBag() *TileBag {
	return &TileBag{
		Tiles: make([]*Tile, 0),
	}
}

func (t *TileBag) Key() string {
	return TileBagKey
}

func (t *TileBag) Diff(other *TileBag) Diff {
	diff := NewDiff()
	if len(t.Tiles) != len(other.Tiles) {
		diff.SetChanged("Tiles", other.Tiles, t.Tiles)
	}

	return diff
}

func (t *TileBag) Clone() state.Model {
	return &TileBag{
		Tiles: t.Tiles,
	}
}
