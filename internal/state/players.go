package state

import (
	"fmt"

	"github.com/konapun/statekit/state"
)

const (
	PlayersKey = "player"
)

var (
	ErrPlayerNotFound = fmt.Errorf("player not found")
)

type Player struct {
	Score int
	Hand  []*Tile
}

func (p *Player) Clone() *Player {
	clone := Player{
		Score: p.Score,
	}
	clone.Hand = make([]*Tile, len(p.Hand))
	for i, tile := range p.Hand {
		clone.Hand[i] = &Tile{
			Shape: tile.Shape,
			Color: tile.Color,
		}
	}
	return &clone
}

type Players struct {
	ActivePlayerIndex int
	Players           []*Player
}

func NewPlayers() *Players {
	return &Players{}
}

func (p *Players) GetActivePlayer() (*Player, error) {
	if len(p.Players) == 0 {
		return nil, ErrPlayerNotFound
	}
	return p.Players[p.ActivePlayerIndex], nil
}

func (p *Players) Key() string {
	return PlayersKey
}

func (p *Players) Diff(other *Players) Diff {
	diff := NewDiff()
	if p.ActivePlayerIndex != other.ActivePlayerIndex {
		diff.SetChanged("ActivePlayerIndex", other.ActivePlayerIndex, p.ActivePlayerIndex)
	}
	if len(p.Players) != len(other.Players) {
		diff.SetChanged("Players", other.Players, p.Players)
	}

	// Check for individual player changes
	playersChanged := false
	for i, player := range p.Players {
		otherPlayer := other.Players[i]
		if player.Score != otherPlayer.Score {
			playersChanged = true
			diff.SetChanged(fmt.Sprintf("Players[%d].Score", i), otherPlayer.Score, player.Score)
		}
		if len(player.Hand) != len(otherPlayer.Hand) {
			playersChanged = true
			diff.SetChanged(fmt.Sprintf("Players[%d].Hand", i), otherPlayer.Hand, player.Hand)
		} else {
			playerHandChanged := false
			for j, tile := range player.Hand {
				if !tile.Equals(otherPlayer.Hand[j]) {
					playerHandChanged = true
					diff.SetChanged(fmt.Sprintf("Players[%d].Hand[%d]", i, j), otherPlayer.Hand[j], tile)
				}
			}
			if playerHandChanged {
				diff.SetChanged(fmt.Sprintf("Players[%d].Hand", i), otherPlayer.Hand, player.Hand)
			}
		}
	}
	if playersChanged {
		diff.SetChanged("Players", other.Players, p.Players)
	}

	return diff
}

func (p *Players) Clone() state.Model {
	clone := &Players{
		ActivePlayerIndex: p.ActivePlayerIndex,
		Players:           make([]*Player, len(p.Players)),
	}
	for i, player := range p.Players {
		clone.Players[i] = player.Clone()
	}
	return clone
}
