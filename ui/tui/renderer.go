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
	fmt.Println("Rendering players", len(players.Players))
	for _, player := range players.Players {
		fmt.Println(player)
	}
	return nil
}

func (r *Renderer) RenderBoard(board *state.Board) error {
	fmt.Println("Rendering board")
	for _, tile := range board.Tiles {
		fmt.Println(tile)
	}
	return nil
}

func (r *Renderer) RenderTileBag(tileBag *state.TileBag) error {
	fmt.Println("Rendering tile bag")
	return nil
}
