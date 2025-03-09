package state_test

import (
	"testing"

	. "github.com/konapun/quirkle/internal/state"
	"github.com/stretchr/testify/require"
)

func TestAccessor_Observers(t *testing.T) {
	accessor := NewAccessor(&TestModel{"test"})

  called := false
	observer1 := NewRuntimeObserver(func(new Model, old Model) {
    called = true
  })

	accessor.RegisterObserver(observer1)
  require.False(t, called)

  accessor.Update(func(m TestModel) error {
    return nil
  })
  require.True(t, called)
}
