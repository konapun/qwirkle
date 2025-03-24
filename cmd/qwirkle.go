package main

import (
	"github.com/konapun/qwirkle/game"
	"github.com/konapun/qwirkle/ui/tui"
)

// Reference implementation
func main() {
	game := game.New(tui.NewInput(), tui.NewObserver())
	game.Run()
}
