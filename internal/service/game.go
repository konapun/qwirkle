package service

import (
	qs "github.com/konapun/qwirkle/internal/state"
)

type GameService struct {
	playersService *PlayersService
	boardService   *BoardService
	tileBagService *TileBagService
}

// TODO: Subscribe to state changes and emit them all at once
type gameObserver struct {
}

func NewGameService(stateManager *qs.Manager) *GameService {
	playersService := NewPlayersService(stateManager.PlayersAccessor)
	boardService := NewBoardService(stateManager.BoardAccessor)
	tileBagService := NewTileBagService(stateManager.TileBagAccessor)

	return &GameService{playersService, boardService, tileBagService}
}

func (g *GameService) AddPlayer() error {
	return g.playersService.AddPlayer()
}

func (g *GameService) GetNumberOfPlayers() int {
	return g.playersService.GetNumberOfPlayers()
}

func (g *GameService) FillTileBag(tiles []*qs.Tile) error {
	return g.tileBagService.AddTiles(tiles)
}

// Place tile takes a tile from the active player's hand and places it on the board
func (g *GameService) PlaceTile(tile *qs.Tile, x, y int) error {
	// Remove the tile from the player's hand
	if err := g.playersService.PlayTile(tile); err != nil {
		return err
	}

	score, err := g.boardService.PlaceTile(tile, x, y)
	if err != nil {
		return err
	}
	g.playersService.IncrementScore(score)
	return nil
}

// DrawTile takes a tile from the tile bag and adds it to the active player's hand
func (g *GameService) DrawTile() error {
	tile, err := g.tileBagService.DrawTile()
	if err != nil {
		return err
	}

	if err = g.playersService.DrawTile(tile); err != nil {
		return err
	}
	return nil
}

func (g *GameService) NextPlayer() error {
	return g.playersService.NextPlayer()
}

func (g *GameService) IsRunning() (bool, error) {
	remainingBaggedTiles := g.tileBagService.GetTiles()
	remainingPlayerTiles, err := g.playersService.GetPlayerHand()
	if err != nil {
		return false, err
	}
	return len(remainingBaggedTiles)+len(remainingPlayerTiles) > 0, nil
}
