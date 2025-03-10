package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestPlayer_Key(t *testing.T) {
	player := Player{}
	require.Equal(t, "player", player.Key())
}

func TestPlayer_Diff(t *testing.T) {
	player := Player{IsActive: true}
	other := Player{IsActive: false, Score: 10}
	diff := player.Diff(&other)
	require.True(t, diff.HasChanged("IsActive"))
	require.False(t, diff.HasChanged("Score"))
}

func TestPlayer_Clone(t *testing.T) {
	player := Player{IsActive: true, Score: 10, Hand: []Tile{{Color: ColorRed, Shape: ShapeCircle}}}
	clone := player.Clone()
	require.Equal(t, player.IsActive, clone.IsActive)
	require.Equal(t, player.Score, clone.Score)
	require.Equal(t, player.Hand, clone.Hand)
}
