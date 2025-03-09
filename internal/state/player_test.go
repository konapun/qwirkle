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
