package main

import (
	"github.com/konapun/qwirkle/ui/tui"
	"github.com/konapun/qwirkle/game"
)

// Reference implementation
func main() {
  // Create the input device
  input := tui.NewInput()

  // Create the output device
  // TODO:

  // Wire up the scene controller
  game := game.New(input)
  game.Run()
}
