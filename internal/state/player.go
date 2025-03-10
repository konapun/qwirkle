package state

const (
	PlayerKey = "player"
)

type Player struct {
	// whether or not this is the active player
	IsActive bool
	Score    int
	Hand     []Tile
}

func (p *Player) Key() string {
	return PlayerKey
}

func (p *Player) Diff(other *Player) Diff {
	diff := NewDiff()
	if p.IsActive != other.IsActive {
		diff.SetChanged("IsActive", other.IsActive, p.IsActive)
	}
	return diff
}

func (p *Player) Clone() *Player {
	return &Player{
		IsActive: p.IsActive,
		Score:    p.Score,
		Hand:     p.Hand,
	}
}
