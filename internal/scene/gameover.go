package scene

import (
	"github.com/konapun/qwirkle/internal"
)

const (
	SceneGameOver = "gameOver"
)

type GameOverAction int

const (
	NewGame GameOverAction = iota
	Quit
)

type GameOver struct {
	input internal.Input[GameOverAction]
}

func NewGameOver(input internal.Input[GameOverAction]) *GameOver {
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
		// start new game
	case Quit:
		// end game
	default:
		return ErrInvalidAction
	}
	return nil
}
