package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestState_Get(t *testing.T) {
	state := NewState(NewTestModel("test1"), NewTestModel("test2"))

	item, err := state.Get("test1")
	require.NoError(t, err)
	require.Equal(t, "test1", item.Key())

	item, err = state.Get("test2")
	require.NoError(t, err)
	require.Equal(t, "test2", item.Key())

	_, err = state.Get("unknown")
	require.Error(t, err)
}
