package main

import (
	"flag"
	"github.com/markoczy/gogameoflife/app"
	"github.com/markoczy/gogameoflife/game"
	"github.com/markoczy/gogameoflife/game/seeder"
	"io/ioutil"
	"log"
	"time"
)

var (
	// Panel dimensions in pixel
	defaultWidth  = 600
	defaultHeight = 600
	// Pixel scaling (does not affect dimension)
	defaultScale = 4
	// Timeout between ticks in milliseconds
	defaultFps = 20
	// Default Seeder
	defaultSeeder  = "random.default"
	defaultDensity = 0.1
)

func main() {
	widthPtr := flag.Int("width", defaultWidth, "width of the panel")
	heightPtr := flag.Int("height", defaultHeight, "height of the panel")
	scalePtr := flag.Int("scale", defaultScale, "pixel scaling")
	fpsPtr := flag.Int("fps", defaultFps, "frames per second")
	seederPtr := flag.String("seeder", defaultSeeder, "seed method")
	densityPtr := flag.Float64("density", defaultDensity,
		"seed density, if applicable")
	debugPtr := flag.Bool("debug", false, "enable debug logs")
	flag.Parse()

	if !(*debugPtr) {
		disableLogs()
	}

	var seed seeder.Seeder
	switch *seederPtr {
	case defaultSeeder:
		nano := time.Now().UTC().UnixNano()
		seed = seeder.CreateDefaultRandom(nano, *densityPtr)
	case "random.gliders":
		nano := time.Now().UTC().UnixNano()
		seed = seeder.CreatePrefabRandom("glider", nano, *densityPtr)
	default:
		log.Panicln("Unrecognized seeder")
		return
	}

	g, err := game.NewGame((*widthPtr)/(*scalePtr), (*heightPtr)/(*scalePtr), seed)
	if err != nil {
		log.Panic(err)
	}
	tick := 1.0 / float64((*fpsPtr)) * 1000.0
	runner := app.NewScreenApp(*widthPtr, *heightPtr, *scalePtr, int(tick))
	err = runner.Run(g)
	if err != nil {
		log.Panic(err)
	}
}

func disableLogs() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func gridToString(grid [][]bool) string {
	ret := ""
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				ret += "+"
			} else {
				ret += "-"
			}
		}
		ret += "\n"
	}
	return ret
}
