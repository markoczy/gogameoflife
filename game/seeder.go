package game

import (
	"log"
	"math/rand"
)

// Seeder needed to create a new game
type Seeder interface {
	Seed(width, height int) (Grid, error)
}

// NewDefaultRandomSeeder creates a new default random seeder (all random)
func NewDefaultRandomSeeder(seed int64, percentage float64) Seeder {
	return &defaultRandom{
		seed:       seed,
		percentage: percentage}
}

// NewPrefabRandomSeeder creates a seeder that creates a prefab at random
func NewPrefabRandomSeeder(prefab string, seed int64, percentage float64) Seeder {
	return &randomSinglePrefab{
		prefab:     prefab,
		seed:       seed,
		percentage: percentage}
}

type defaultRandom struct {
	percentage float64
	seed       int64
}

func (d *defaultRandom) Seed(width, height int) (Grid, error) {
	rand.Seed(int64(d.seed))
	data := make([][]bool, height)
	for iRow := range data {
		data[iRow] = make([]bool, width)
		for iCell := range data[iRow] {
			data[iRow][iCell] = rand.Float64() > (1 - d.percentage)
		}
	}
	return Grid{
		Data:   data,
		Width:  width,
		Height: height}, nil
}

type randomSinglePrefab struct {
	prefab     string
	percentage float64
	seed       int64
}

func (d *randomSinglePrefab) Seed(width, height int) (Grid, error) {
	rand.Seed(int64(d.seed))
	grid := NewEmptyGrid(width, height)
	prefab := GetPrefab(d.prefab, 0, false, false)

	count := int((float64(width) * float64(height) * (d.percentage * d.percentage)) / float64(prefab.Width*prefab.Height))
	log.Println("Spawn Count:", count)

	for i := 0; i < count; i++ {
		x := int(rand.Float64() * float64(width))
		y := int(rand.Float64() * float64(height))
		rot := int(rand.Float64()*256.0) % 4
		flipX := rand.Float64() > 0.5
		flipY := rand.Float64() > 0.5
		prefab = GetPrefab(d.prefab, rot, flipX, flipY)
		log.Printf("Seeding x: %d, y: %d, rot: %d, flipX: %v, flipY: %v\n", x, y, rot, flipX, flipY)
		grid = grid.InsertGridAt(prefab, x, y)
	}
	return grid, nil
}
