package service_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/service"
	qs "github.com/konapun/qwirkle/internal/state"
	"github.com/konapun/statekit/state"
	"github.com/stretchr/testify/require"
)

func getPlayersService(players qs.Players, observers ...state.Observer[*qs.Players]) *PlayersService {
	accessor := state.NewAccessor(&players)
	service := NewPlayersService(accessor)
	for _, observer := range observers {
		accessor.RegisterObserver(observer)
	}

	return service
}

func TestPlayerService_PlayTile(t *testing.T) {
	redCircle := qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	player := qs.Player{
		Hand: []*qs.Tile{&redCircle},
	}
	playersService := getPlayersService(qs.Players{Players: []*qs.Player{&player}})

	// Test playing a tile that doesn't exist in the player's hand
	err := playersService.PlayTile(&qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeSquare})
	require.Equal(t, ErrTileNotFound, err)

	// Test playing a tile that exists in the player's hand
	err = playersService.PlayTile(&redCircle)
	require.Nil(t, err)

	// Test that the tile was removed from the player's hand
	err = playersService.PlayTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Equal(t, ErrTileNotFound, err)
}

func TestPlayerService_NextPlayer(t *testing.T) {
	player1 := qs.Player{}
	player2 := qs.Player{}

	observerCalled := false
	observer := state.NewRuntimeObserver(func(new *qs.Players, old *qs.Players) {
		observerCalled = true
		require.NotEqual(t, old.ActivePlayerIndex, new.ActivePlayerIndex)
		require.Equal(t, 0, old.ActivePlayerIndex)
		require.Equal(t, 1, new.ActivePlayerIndex)
	})
	playersService := getPlayersService(qs.Players{Players: []*qs.Player{&player1, &player2}}, observer)

	playersService.NextPlayer()
	require.True(t, observerCalled)
}

func TestPlayerService_HasTile(t *testing.T) {
	redCircle := qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	blueSquare := qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeSquare}
	player := qs.Player{
		Hand: []*qs.Tile{&redCircle},
	}
	playerService := getPlayersService(qs.Players{Players: []*qs.Player{&player}})

	// Test checking for a tile that exists in the player's hand
	hasTile, err := playerService.HasTile(&redCircle)
	require.NoError(t, err)
	require.True(t, hasTile)

	// Test checking for a tile that doesn't exist in the player's hand
	hasTile, err = playerService.HasTile(&blueSquare)
	require.NoError(t, err)
	require.False(t, hasTile)
}

func TestPlayerService_DrawTile(t *testing.T) {
	player := qs.Player{}
	playerService := getPlayersService(qs.Players{Players: []*qs.Player{&player}})

	// Test drawing a tile when the hand is not full
	err := playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Nil(t, err)

	// Test drawing a tile when the hand is full
	err = playerService.DrawTile(&qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle})
	require.Equal(t, ErrHandFull, err)
}

func TestPlayerService_IncrementScore(t *testing.T) {
	player := qs.Player{}
	playersService := getPlayersService(qs.Players{Players: []*qs.Player{&player}})

	// Test incrementing the score by a positive amount
	err := playersService.IncrementScore(10)
	require.Nil(t, err)

	// Test incrementing the score by a negative amount
	err = playersService.IncrementScore(-10)
	require.Equal(t, ErrInvalidScore, err)
}

func TestPlayerService_GetPlayerHand(t *testing.T) {
	redCircle := qs.Tile{Color: qs.ColorRed, Shape: qs.ShapeCircle}
	blueSquare := qs.Tile{Color: qs.ColorBlue, Shape: qs.ShapeSquare}
	orangeClover := qs.Tile{Color: qs.ColorOrange, Shape: qs.ShapeClover}
	player := qs.Player{
		Hand: []*qs.Tile{&redCircle, &blueSquare, &orangeClover},
	}

	playersService := getPlayersService(qs.Players{Players: []*qs.Player{&player}})
	hand, err := playersService.GetPlayerHand()
  require.NoError(t, err)
	require.Len(t, hand, 3)
	require.Equal(t, &redCircle, hand[0])
	require.Equal(t, &blueSquare, hand[1])
	require.Equal(t, &orangeClover, hand[2])
}
