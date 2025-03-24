package tui

import (
	"fmt"

	"github.com/konapun/qwirkle/internal/state"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderPlayers(players *state.Players) error {
	fmt.Println("Rendering players")
	return nil
}

func (r *Renderer) RenderBoard(board *state.Board) error {
	fmt.Println("Rendering board")
	return nil
}

func (r *Renderer) RenderTileBag(tileBag *state.TileBag) error {
	fmt.Println("Rendering tile bag")
	return nil
}
