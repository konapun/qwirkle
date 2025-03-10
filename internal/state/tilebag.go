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
	ColorRed Color = iota
	ColorOrange
	ColorYellow
	ColorGreen
	ColorBlue
	ColorPurple
)

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
	Tiles []Tile
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

func (t *TileBag) Clone() *TileBag {
	return &TileBag{
		Tiles: t.Tiles,
	}
}
