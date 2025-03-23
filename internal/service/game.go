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

// Place tile takes a tile run from the active player's hand and places it on the board
func (g *GameService) PlaceTiles(tileRun *qs.Run) error {
	// Ensure the player has the tiles in the run and remove them from the player's hand
	return g.playersService.PlayTiles(tileRun.Tiles, func() error {
		score, err := g.boardService.PlaceTiles(tileRun)
		if err != nil {
			return err
		}
		g.playersService.IncrementScore(score)
		return nil
	})
}

func (g *GameService) SwapTiles(tiles []*qs.Tile) error {
	return g.playersService.PlayTiles(tiles, func() error {
		// Add the tiles back to the bag
		if err := g.tileBagService.AddTiles(tiles); err != nil {
			return err
		}
		// Draw new tiles to replace the ones swapped
		for range tiles {
			tile, err := g.tileBagService.DrawTile()
			if err != nil {
				return err
			}
			if err = g.playersService.DrawTile(tile); err != nil {
				return err
			}
		}
		return nil
	})
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

func (g *GameService) IsOver() (bool, error) {
	remainingBaggedTiles := g.tileBagService.GetTiles()
	remainingPlayerTiles, err := g.playersService.GetPlayerHand()
	if err != nil {
		return false, err
	}
	return len(remainingBaggedTiles)+len(remainingPlayerTiles) > 0, nil
}
