package seeder

import (
	"github.com/markoczy/gogameoflife/game/grid"
	"math/rand"
)

type defaultRandom struct {
	percentage float64
	seed       int64
}

func (d *defaultRandom) Seed(width, height int) (grid.Grid, error) {
	rand.Seed(int64(d.seed))
	data := make([][]bool, height)
	for iRow := range data {
		data[iRow] = make([]bool, width)
		for iCell := range data[iRow] {
			data[iRow][iCell] = rand.Float64() > (1 - d.percentage)
		}
	}
	return grid.Grid{
		Data:   data,
		Width:  width,
		Height: height}, nil
}
