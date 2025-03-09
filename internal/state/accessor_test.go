package state_test

import (
	"errors"
	"testing"

	. "github.com/konapun/qwirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestAccessor_Observers(t *testing.T) {
	testModel := NewTestModel("test")
	testModel.String = "initial"
	accessor := NewAccessor(testModel)

	called := false
	var newModel, oldModel *TestModel
	observer1 := NewRuntimeObserver(func(new *TestModel, old *TestModel) {
		called = true
		newModel = new
		oldModel = old
	})

	accessor.RegisterObserver(observer1)
	require.False(t, called)

	accessor.Update(func(m *TestModel) error {
		m.String = "updated"
		return nil
	})
	require.True(t, called)
	require.Equal(t, "initial", oldModel.String)
	require.Equal(t, "updated", newModel.String)

	// Test updating the model with an error
	called = false
	err := accessor.Update(func(m *TestModel) error {
		return errors.New("error")
	})
	require.Error(t, err)
	require.False(t, called)
}
