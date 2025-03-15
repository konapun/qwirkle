package scene_test

import (
	"errors"
	"testing"

	. "github.com/konapun/qwirkle/internal/scene"
	"github.com/stretchr/testify/require"
)

func TestController_Transition(t *testing.T) {
	scene1 := NewScene("scene1", func(controller *Controller) error {
		return nil
	})
	scene2 := NewScene("scene2", func(controller *Controller) error {
		return errors.New("scene2 error")
	})
	controller := NewController(scene1, scene2)
	err := controller.Transition("scene1")
	require.NoError(t, err)
	err = controller.Transition("scene2")
	require.Error(t, err)
	err = controller.Transition("unknown")
	require.Error(t, err)
	require.Equal(t, ErrInvalidTransition, err)
}
