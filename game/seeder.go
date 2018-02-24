package game

import (
	"math/rand"
)

// Seeder needed to create a new game
type Seeder interface {
	Seed(g Game) ([][]bool, error)
}

// Creates a default random seeder (all random)
func NewDefaultRandomSeeder(seed int64, percentage float64) Seeder {
	return &defaultRandom {
		seed: seed,
		percentage: percentage }
}

type defaultRandom struct {
	percentage float64
	seed int64
}

func (d *defaultRandom) Seed(g Game) ([][]bool, error) {
	rand.Seed(int64(d.seed))
	data := make([][]bool, g.Height())
	for iRow, _ := range data {
		data[iRow] = make([]bool, g.Width())
		for iCell, _ := range data[iRow] {
			data[iRow][iCell] = rand.Float64() > (1 - d.percentage)
		}
	}
	return data, nil
}