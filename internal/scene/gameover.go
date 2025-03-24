package scene

import (
	"github.com/konapun/qwirkle/internal/io"
)

const SceneGameOver = "gameOver"

const (
	NewGame GameOverAction = iota
	Quit
)

type GameOverAction int

type GameOver struct {
	input io.Input[GameOverAction]
}

func NewGameOver(input io.Input[GameOverAction]) *GameOver {
	return &GameOver{
		input: input,
	}
}

func (g *GameOver) Key() string {
	return SceneGameOver
}

func (g *GameOver) Run(controller *Controller) error {
	action := g.input.Read()
	switch action {
	case NewGame:
		controller.Transition(SceneStartGame)
	case Quit:
		return nil
	default:
		return ErrInvalidAction
	}
	return nil
}
