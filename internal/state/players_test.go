package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestPlayers_Key(t *testing.T) {
	players := Players{}
	require.Equal(t, "player", players.Key())
}

func TestPlayers_Diff(t *testing.T) {
	players := Players{
		ActivePlayerIndex: 0,
		Players: []*Player{{
			Score: 0,
			Hand:  []*Tile{{Color: ColorRed, Shape: ShapeCircle}},
		}},
	}
	other := Players{
		ActivePlayerIndex: 1,
		Players: []*Player{{
			Score: 1,
			Hand:  []*Tile{{Color: ColorBlue, Shape: ShapeCircle}},
		}},
	}

	diff := players.Diff(&other)
	require.True(t, diff.HasChanged("ActivePlayerIndex"))
	require.True(t, diff.HasChanged("Players"))
	require.True(t, diff.HasChanged("Players[0].Score"))
	require.True(t, diff.HasChanged("Players[0].Hand"))
	require.True(t, diff.HasChanged("Players[0].Hand[0]"))
}

func TestPlayers_Clone(t *testing.T) {
	player1 := Player{Score: 0, Hand: []*Tile{{Color: ColorRed, Shape: ShapeCircle}}}
	player2 := Player{Score: 1, Hand: []*Tile{{Color: ColorBlue, Shape: ShapeCircle}}}
	players := Players{Players: []*Player{&player1, &player2}}
	clone := players.Clone().(*Players)
	require.Len(t, clone.Players, 2)
	for i := range clone.Players {
		require.Equal(t, players.Players[i].Score, clone.Players[i].Score)
		for j := range clone.Players[i].Hand {
			require.True(t, players.Players[i].Hand[j].Equals(clone.Players[i].Hand[j]))
		}
	}
}

func TestPlayers_GetActivePlayer(t *testing.T) {

}
