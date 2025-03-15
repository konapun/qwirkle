package service

import (
	"errors"

	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
)

var (
	ErrNoTiles        = errors.New("no tiles left in the bag")
	ErrNotEnoughTiles = errors.New("not enough tiles in the bag")
)

type TileBagService struct {
	accessor *state.Accessor[*qs.TileBag]
}

func NewTileBagService(accessor *state.Accessor[*qs.TileBag]) *TileBagService {
	return &TileBagService{accessor}
}

// AddTiles adds the given tiles to the bag
func (t *TileBagService) AddTiles(tiles []*qs.Tile) error {
	return t.accessor.Update(func(b *qs.TileBag) error {
		b.Tiles = append(b.Tiles, tiles...)
		return nil
	})
}

// DrawTile removes a tile from the bag and returns it
func (t *TileBagService) DrawTile() (*qs.Tile, error) {
	var tile *qs.Tile
	if err := t.accessor.Update(func(b *qs.TileBag) error {
		if len(b.Tiles) == 0 {
			return ErrNoTiles
		}
		// Draw a tile from the bag
		tile = b.Tiles[len(b.Tiles)-1]
		b.Tiles = b.Tiles[:len(b.Tiles)-1]
		return nil
	}); err != nil {
		return nil, err
	}
	return tile, nil
}

// ExchangeTiles swaps the given tiles for new ones, adding the old ones back to the bag
func (t *TileBagService) ExchangeTiles(tiles []*qs.Tile) ([]*qs.Tile, error) {
	var newTiles []*qs.Tile
	if err := t.accessor.Update(func(b *qs.TileBag) error {
		// Draw new tiles from the bag
		for range tiles {
			if len(b.Tiles) == 0 {
				return ErrNotEnoughTiles
			}
			newTile := b.Tiles[len(b.Tiles)-1]
			b.Tiles = b.Tiles[:len(b.Tiles)-1]
			newTiles = append(newTiles, newTile)
		}

		// Add the old tiles back to the bag
		b.Tiles = append(b.Tiles, tiles...)

		return nil
	}); err != nil {
		return nil, err
	}
	return newTiles, nil
}

func (t *TileBagService) GetTiles() []*qs.Tile {
	return t.accessor.Query().Tiles
}
