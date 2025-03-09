package service

import (
	"errors"
	"slices"

	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

const (
	MaxHandSize = 6
)

var (
	ErrTileNotFound = errors.New("tile not found")
	ErrHandFull     = errors.New("hand is full")
	ErrInvalidScore = errors.New("score must be positive")
)

type PlayerService struct {
	accessor *state.Accessor[*qs.Player]
}

func NewPlayerService(accessor *state.Accessor[*qs.Player]) *PlayerService {
	return &PlayerService{accessor}
}

func (p *PlayerService) PlayTile(tile *qs.Tile) error {
	return p.accessor.Update(func(p *qs.Player) error {
		for i, t := range p.Hand {
			if t == *tile {
				p.Hand = slices.Delete(p.Hand, i, i+1)
				return nil
			}
		}
		return ErrTileNotFound
	})
}

// SetActive sets the player's active status
func (p *PlayerService) SetActive(active bool) error {
	return p.accessor.Update(func(b *qs.Player) error {
		b.IsActive = active
		return nil
	})
}

// DrawTile places a tile into the player's hand
func (p *PlayerService) DrawTile(tile *qs.Tile) error {
	return p.accessor.Update(func(p *qs.Player) error {
		if len(p.Hand) == MaxHandSize {
			return ErrHandFull
		}
		p.Hand = append(p.Hand, *tile)
		return nil
	})
}

// IncrementScore adds the given amount to the player's score
func (p *PlayerService) IncrementScore(amount int) error {
	return p.accessor.Update(func(b *qs.Player) error {
		if amount < 0 {
			return ErrInvalidScore
		}
		b.Score += amount
		return nil
	})
}
