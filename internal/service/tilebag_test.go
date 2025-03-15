package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
	"github.com/stretchr/testify/require"
)

func TestTileBagService_AddTiles(t *testing.T) {
	tileBag := qs.TileBag{}
	accessor := state.NewAccessor(&tileBag)
	tileBagService := NewTileBagService(accessor)

	// Test adding a single tile
	err := tileBagService.AddTiles([]*qs.Tile{{Color: qs.ColorRed, Shape: qs.ShapeCircle}})
	require.Nil(t, err)
	require.Len(t, tileBag.Tiles, 1)
}

func TestTileBagService_DrawTile(t *testing.T) {
	tileBag := qs.TileBag{Tiles: []*qs.Tile{{Color: qs.ColorRed, Shape: qs.ShapeCircle}}}
	accessor := state.NewAccessor(&tileBag)
	tileBagService := NewTileBagService(accessor)

	// Test drawing a tile when the bag is not empty
	tile, err := tileBagService.DrawTile()
	require.Nil(t, err)
	require.Len(t, tileBag.Tiles, 0)
	require.Equal(t, qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}, *tile)

	// Test drawing a tile when the bag is empty
	_, err = tileBagService.DrawTile()
	require.Equal(t, ErrNoTiles, err)
}

func TestTileBagService_ExchangeTiles(t *testing.T) {
	redCircleTile := &qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	greenSquareTile := &qs.Tile{Color: qs.ColorGreen, Shape: qs.ShapeSquare}
	orangeDiamondTile := &qs.Tile{Color: qs.ColorOrange, Shape: qs.ShapeDiamond}
	tileBag := qs.TileBag{Tiles: []*qs.Tile{redCircleTile}}
	accessor := state.NewAccessor(&tileBag)
	tileBagService := NewTileBagService(accessor)

	// Test exchanging a single tile
	newTiles, err := tileBagService.ExchangeTiles([]*qs.Tile{greenSquareTile})
	require.Nil(t, err)
	require.Len(t, tileBag.Tiles, 1)
	require.Len(t, newTiles, 1)
	require.Equal(t, greenSquareTile, accessor.Query().Tiles[0])
	require.Equal(t, redCircleTile, newTiles[0])

	// Test trying to exchange more tiles than are in the bag
	_, err = tileBagService.ExchangeTiles([]*qs.Tile{redCircleTile, orangeDiamondTile})
	require.Equal(t, ErrNotEnoughTiles, err)
}
