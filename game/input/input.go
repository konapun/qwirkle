package input

import (
	"strings"

	"github.com/konapun/qwirkle/internal/scene"
)

type StringReader interface {
	Read() (string, error)
}

// Input translates text-based user input into actions.
type Input struct {
	StartGameReader  StartGameReader
	PlayerTurnReader PlayerTurnReader
	GameOverReader   GameOverReader
}

func NewInput(reader StringReader) *Input {
	return &Input{
		StartGameReader:  StartGameReader{reader},
		PlayerTurnReader: PlayerTurnReader{reader},
		GameOverReader:   GameOverReader{reader},
	}
}

type StartGameReader struct {
	reader StringReader
}

func (g *StartGameReader) Read() scene.StartGameAction {
	reader := g.reader
	// FIXME: handle err
	input, _ := reader.Read()
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "add":
		return scene.AddPlayer
	case "start":
		return scene.Start
	default:
		return scene.StartGameAction(scene.SceneActionUnknown)
	}
}

type PlayerTurnReader struct {
	reader StringReader
}

func (g *PlayerTurnReader) Read() scene.PlayerAction {
	return scene.PlayerAction{Type: scene.PlaceTiles}
}

type GameOverReader struct {
	reader StringReader
}

func (g *GameOverReader) Read() scene.GameOverAction {
	return scene.GameOverAction(1)
}
