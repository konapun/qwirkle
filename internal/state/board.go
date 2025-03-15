package state

import (
	"errors"
	"sort"

	"github.com/konapun/statekit/state"
)

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

var (
	ErrCellOccupied = errors.New("cell occupied")
)

type Direction int

type Type int

type Line []*Cell

type Cell struct {
	X, Y int
	Tile *Tile
}

func (l Line) Length() int {
	return len(l)
}

func (l Line) Contains(tile *Tile) bool {
	for _, cell := range l {
		if cell.Tile.Equals(tile) {
			return true
		}
	}
	return false
}

func (l Line) Type() Type {
	if l.Length() < 2 {
		return TypeUndetermined
	}
	if l[0].Tile.Color == l[1].Tile.Color {
		return TypeColorMatch
	}
	if l[0].Tile.Shape == l[1].Tile.Shape {
		return TypeShapeMatch
	}
	// Shouldn't happen
	return TypeUndetermined
}

func (l Line) IsValid() bool {
	switch l.Type() {
	case TypeUndetermined:
		return l.Length() == 1
	case TypeColorMatch:
		shapes := make(map[Shape]bool)
		for _, cell := range l {
			if shapes[cell.Tile.Shape] {
				return false
			}
			shapes[cell.Tile.Shape] = true
		}
		return true
	case TypeShapeMatch:
		colors := make(map[Color]bool)
		for _, cell := range l {
			if colors[cell.Tile.Color] {
				return false
			}
			colors[cell.Tile.Color] = true
		}
		return true
	}
	return false
}

func (l Line) GetTiles() []*Tile {
	tiles := make([]*Tile, 0)
	for _, cell := range l {
		tiles = append(tiles, cell.Tile)
	}
	return tiles
}

func (l Line) GetShapes() []Shape {
	shapes := make([]Shape, 0)
	for _, cell := range l {
		shapes = append(shapes, cell.Tile.Shape)
	}
	return shapes
}

func (l Line) GetColors() []Color {
	colors := make([]Color, 0)
	for _, cell := range l {
		colors = append(colors, cell.Tile.Color)
	}
	return colors
}

// Sort sorts the line by Y, then X to ensure proper ordering while reading line contents
func (l Line) Sort() Line {
	sortedLine := make(Line, len(l))
	copy(sortedLine, l)

	sort.Slice(sortedLine, func(i, j int) bool {
		if sortedLine[i].Y == sortedLine[j].Y {
			return sortedLine[i].X < sortedLine[j].X
		}
		return sortedLine[i].Y < sortedLine[j].Y
	})
	return sortedLine
}

type Board struct {
	Tiles map[[2]int]*Tile
}

// NewBoard creates a new board with an empty map of tiles
func NewBoard() *Board {
	return &Board{
		Tiles: make(map[[2]int]*Tile),
	}
}

// GetLine returns a line of tiles starting at the given position in the given direction until it encounters an empty space
func (b *Board) GetLine(x, y int, direction Direction) Line {
	cells := make([]*Cell, 0)
	line := Line(cells)
	collectCells := func(x, y int, dx, dy int) {
		for {
			tile, exists := b.Tiles[[2]int{x, y}]
			if !exists {
				break
			}
			line = append(line, &Cell{X: x, Y: y, Tile: tile})
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

	return line.Sort()
}

// Test returns the lines that would be formed by placing a tile at the given position
func (b *Board) Test(tile *Tile, x, y int) (Line, Line, error) {
	if _, exists := b.Tiles[[2]int{x, y}]; exists {
		return nil, nil, ErrCellOccupied
	}

	b.Tiles[[2]int{x, y}] = tile
	horizontal := b.GetLine(x, y, DirectionHorizontal)
	vertical := b.GetLine(x, y, DirectionVertical)
	delete(b.Tiles, [2]int{x, y})

	return horizontal, vertical, nil
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

func (b *Board) Clone() state.Model {
	newTiles := make(map[[2]int]*Tile)
	for k, v := range b.Tiles {
		newTiles[k] = v
	}
	return &Board{
		Tiles: newTiles,
	}
}
