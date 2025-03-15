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
	ErrPlayerNotFound = errors.New("player not found")
	ErrTileNotFound   = errors.New("tile not found")
	ErrHandFull       = errors.New("hand is full")
	ErrInvalidScore   = errors.New("score must be positive")
)

type PlayersService struct {
	accessor *state.Accessor[*qs.Players]
}

func NewPlayersService(accessor *state.Accessor[*qs.Players]) *PlayersService {
	return &PlayersService{accessor}
}

func (p *PlayersService) AddPlayer() error {
	return p.accessor.Update(func(p *qs.Players) error {
		p.Players = append(p.Players, &qs.Player{})
		return nil
	})
}

func (p *PlayersService) GetNumberOfPlayers() int {
	players := p.accessor.Query()
	return len(players.Players)
}

func (p *PlayersService) NextPlayer() error {
	return p.accessor.Update(func(p *qs.Players) error {
		if len(p.Players) == 0 {
			return ErrPlayerNotFound
		}
		p.ActivePlayerIndex = (p.ActivePlayerIndex + 1) % len(p.Players)
		return nil
	})
}

func (p *PlayersService) HasTile(tile *qs.Tile) (bool, error) {
	players := p.accessor.Query()
	activePlayer, err := players.GetActivePlayer()
	if err != nil {
		return false, ErrPlayerNotFound
	}
	for _, t := range activePlayer.Hand {
		if t.Equals(tile) {
			return true, nil
		}
	}
	return false, nil
}

func (p *PlayersService) PlayTile(tile *qs.Tile) error {
	return p.accessor.Update(func(p *qs.Players) error {
		activePlayer, err := p.GetActivePlayer()
		if err != nil {
			return ErrPlayerNotFound
		}

		for i, t := range activePlayer.Hand {
			if t == tile {
				activePlayer.Hand = slices.Delete(activePlayer.Hand, i, i+1)
				return nil
			}
		}
		return ErrTileNotFound
	})
}

// DrawTile places a tile into the player's hand
func (p *PlayersService) DrawTile(tile *qs.Tile) error {
	return p.accessor.Update(func(p *qs.Players) error {
		activePlayer, err := p.GetActivePlayer()
		if err != nil {
			return ErrPlayerNotFound
		}

		if len(activePlayer.Hand) == MaxHandSize {
			return ErrHandFull
		}
		activePlayer.Hand = append(activePlayer.Hand, tile)
		return nil
	})
}

// IncrementScore adds the given amount to the player's score
func (p *PlayersService) IncrementScore(amount int) error {
	return p.accessor.Update(func(p *qs.Players) error {
		if amount < 0 {
			return ErrInvalidScore
		}

		activePlayer, err := p.GetActivePlayer()
		if err != nil {
			return ErrPlayerNotFound
		}

		activePlayer.Score += amount
		return nil
	})
}

func (p *PlayersService) GetPlayerHand() ([]*qs.Tile, error) {
	players := p.accessor.Query()
	activePlayer, err := players.GetActivePlayer()
	if err != nil {
		return nil, ErrPlayerNotFound
	}
	return activePlayer.Hand, nil
}
