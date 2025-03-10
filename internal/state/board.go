package state

const (
	BoardKey = "board"
)

const (
	DirectionVertical Direction = iota
	DirectionHorizontal
)

const (
	// TypeColorMatch is a line where all tiles have the same color
	TypeColorMatch Type = iota
	// TypeShapeMatch is a line where all tiles have the same shape
	TypeShapeMatch
	// TypeUndetermined is a line which does not have enough tiles to determine its type
	TypeUndetermined
)

type Direction int

type Type int

type Line []*Tile

func (l Line) Length() int {
	return len(l)
}

func (l Line) Contains(tile *Tile) bool {
	for _, t := range l {
		if t.Equals(tile) {
			return true
		}
	}
	return false
}

func (l Line) Type() Type {
	if l.Length() < 2 {
		return TypeUndetermined
	}
	if l[0].Color == l[1].Color {
		return TypeColorMatch
	}
	if l[0].Shape == l[1].Shape {
		return TypeShapeMatch
	}
	// Shouldn't happen
	return TypeUndetermined
}

type Board struct {
	// the tiles on the board
	Tiles [][]*Tile
}

// GetLine returns a line of tiles starting at the given position in the given direction until it encounters an empty space
func (b *Board) GetLine(x, y int, direction Direction) Line {
	line := make([]*Tile, 0)
	collectCells := func(x, y int, dx, dy int) {
		for {
			tile := b.Tiles[x][y]
			if tile == nil {
				break
			}
			line = append(line, tile)
			x += dx
			y += dy
		}
	}

	// Collect cells in the positive direction
	switch direction {
	case DirectionVertical:
		collectCells(x, y, 0, 1)
	case DirectionHorizontal:
		collectCells(x, y, 1, 0)
	}

	// Collect cells in the negative direction (excluding the starting cell)
	switch direction {
	case DirectionVertical:
		collectCells(x, y-1, 0, -1)
	case DirectionHorizontal:
		collectCells(x-1, y, -1, 0)
	}

	return line
}

func (b *Board) Key() string {
	return BoardKey
}

func (b *Board) Diff(other *Board) Diff {
	diff := NewDiff()
	if len(b.Tiles) != len(other.Tiles) {
		diff.SetChanged("Tiles", other.Tiles, b.Tiles)
	}
	return diff
}

func (b *Board) Clone() *Board {
	return &Board{
		Tiles: b.Tiles,
	}
}
