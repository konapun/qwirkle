package state

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
	Red Color = iota
	Orange
	Yellow
	Green
	Blue
	Purple
)

type Shape int

type Color int

type Tile struct {
	Shape Shape
	Color Color
}

type TileBag struct {
	Tiles []Tile
}

func (t *TileBag) Key() string {
	return TileBagKey
}

func (t *TileBag) Diff(other *Tile) Diff {
	return Diff{}
}

func (t *TileBag) Clone() *TileBag {
	return &TileBag{
		Tiles: t.Tiles,
	}
}
