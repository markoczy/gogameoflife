package seeder

import (
	"github.com/markoczy/gogameoflife/game/grid"
	"github.com/markoczy/gogameoflife/game/prefabs"
	"log"
	"math/rand"
)

type randomSinglePrefab struct {
	prefab     string
	percentage float64
	seed       int64
}

func (d *randomSinglePrefab) Seed(width, height int) (*grid.Grid, error) {
	rand.Seed(int64(d.seed))
	grid := grid.CreateEmpty(width, height)
	prefab := prefabs.Get(d.prefab, 0, false, false)

	count := int((float64(width) * float64(height) * (d.percentage * d.percentage)) / float64(prefab.Width*prefab.Height))
	log.Println("Spawn Count:", count)

	for i := 0; i < count; i++ {
		x := int(rand.Float64() * float64(width))
		y := int(rand.Float64() * float64(height))
		rot := int(rand.Float64()*256.0) % 4
		flipX := rand.Float64() > 0.5
		flipY := rand.Float64() > 0.5
		prefab = prefabs.Get(d.prefab, rot, flipX, flipY)
		log.Printf("Seeding x: %d, y: %d, rot: %d, flipX: %v, flipY: %v\n", x, y, rot, flipX, flipY)
		grid.InsertGridAt(prefab, x, y)
	}
	return grid, nil
}
