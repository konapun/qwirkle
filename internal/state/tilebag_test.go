package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestTileBag_Key(t *testing.T) {
	tileBag := TileBag{}
	require.Equal(t, "tileBag", tileBag.Key())
}

func TestTileBag_Diff(t *testing.T) {
	tileBag := TileBag{Tiles: []Tile{{Color: ColorRed, Shape: ShapeCircle}}}
	other := TileBag{Tiles: []Tile{{Color: ColorRed, Shape: ShapeCircle}, {Color: ColorBlue, Shape: ShapeSquare}}}
	diff := tileBag.Diff(&other)
	require.True(t, diff.HasChanged("Tiles"))
}

func TestTileBag_Clone(t *testing.T) {
	tileBag := TileBag{Tiles: []Tile{{Color: ColorRed, Shape: ShapeCircle}}}
	clone := tileBag.Clone()
	require.Equal(t, tileBag.Tiles, clone.Tiles)
}
