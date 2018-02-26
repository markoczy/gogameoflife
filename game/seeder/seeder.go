package seeder

import "github.com/markoczy/gogameoflife/game/grid"

// Seeder needed to create a new game
type Seeder interface {
	Seed(width, height int) (grid.Grid, error)
}

// CreateDefaultRandom creates a new default random seeder (all random)
func CreateDefaultRandom(seed int64, percentage float64) Seeder {
	return &defaultRandom{
		seed:       seed,
		percentage: percentage}
}

// CreatePrefabRandom creates a seeder that creates a prefab at random
func CreatePrefabRandom(prefab string, seed int64, percentage float64) Seeder {
	return &randomSinglePrefab{
		prefab:     prefab,
		seed:       seed,
		percentage: percentage}
}
