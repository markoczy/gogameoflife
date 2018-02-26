package game

import (
	"github.com/markoczy/gogameoflife/game/grid"
	"github.com/markoczy/gogameoflife/game/seeder"
)

// NewGame creates a new game with the provided seeder
func NewGame(width, height int, seeder seeder.Seeder) (Game, error) {
	// refactor me
	ret := game{
		grid: grid.Grid{Width: width, Height: height}}
	data, err := seeder.Seed(width, height)
	if err != nil {
		return nil, err
	}
	ret.grid = *data
	return &ret, nil
}

// Game exported interface of a game
type Game interface {
	Width() int
	Height() int
	Tick() error
	Render() ([][]bool, error)
}

func (g *game) Width() int {
	return g.grid.Width
}

func (g *game) Height() int {
	return g.grid.Height
}

func (g *game) Render() ([][]bool, error) {
	return g.grid.Data, nil
}

func (g *game) Tick() error {
	data := make([][]bool, g.grid.Height)
	for iRow, row := range g.grid.Data {
		data[iRow] = make([]bool, g.grid.Width)
		for iCell := range row {
			neighbours := g.grid.CountNeighboursAt(iCell, iRow)
			if !g.grid.Data[iRow][iCell] {
				if neighbours == 3 {
					data[iRow][iCell] = true // live
				}
			} else {
				if neighbours < 2 || neighbours > 3 {
					data[iRow][iCell] = false // die
				} else {
					data[iRow][iCell] = true
				}
			}
		}
	}
	g.grid.Data = data
	return nil
}

type game struct {
	grid grid.Grid
}

// func (g *game) getCellAt(x, y int) bool {
// 	return g.data[mod(y, g.height)][mod(x, g.width)]
// }

// func mod(val, mod int) int {
// 	if val >= 0 {
// 		return val % mod
// 	}
// 	for val < 0 {
// 		val += mod
// 	}
// 	return val
// }
//
