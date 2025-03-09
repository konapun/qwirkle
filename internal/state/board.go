package state

const (
	BoardKey = "board"
)

type Board struct {
	// the tiles on the board
	Tiles [][]Tile
}

func (b *Board) Key() string {
	return BoardKey
}

func (b *Board) Diff(other *Board) Diff {
	return Diff{}
}

func (b *Board) Clone() *Board {
	return &Board{
		Tiles: b.Tiles,
	}
}
