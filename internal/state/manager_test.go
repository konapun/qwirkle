package state_test

import (
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestManager_AccessorFor(t *testing.T) {
	state := NewState(NewTestModel("test"))
	manager := NewManager(state)

	// Test trying to get an accessor for a bad key which is not in the state
	accessor, err := manager.AccessorFor("unknown")
	require.Nil(t, accessor)
	require.Error(t, err)

	// Test getting a new accessor
	accessor, err = manager.AccessorFor("test")
	require.NoError(t, err)

	// Test getting a cached accessor
	accessor, err = manager.AccessorFor("test")
	require.NoError(t, err)
}
