package app

import (
	"log"
	"time"
	"image"
	"image/color"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"github.com/markoczy/gogameoflife/game")


var (
	black    = color.RGBA{0x00, 0x00, 0x00, 0x00}
	white    = color.RGBA{0xff, 0xff, 0xff, 0xff}
	pos0     = image.Point{0, 0}

	interrupt = false
)

type App interface {
	Run(g game.Game) error
}

func NewScreenApp(width, height, scale, tick int) App {
	return &screenApp {
		width: width,
		height: height,
		scale: scale,
		tick: tick }
}

type screenApp struct {
	width, height, scale int
	tick int
}

func (app *screenApp) Run(g game.Game) error {
	chErr := make(chan error)
	chInterrupt := make(chan bool)
	bounds := image.Point{app.width, app.height}

	// Start Application and Game Thread
	go driver.Main(func(s screen.Screen) {
		// Init Window
		w, err := s.NewWindow(&screen.NewWindowOptions {
			Width: app.width, 
			Height: app.height, 
			Title: "Application Window" })
		if err != nil { 
			chErr <- err
			return
		}

		// Start Game Thread
		go func() {
			for !interrupt {
				tStart := time.Now().UTC().UnixNano()
				bg, err := s.NewTexture(bounds)
				if err != nil {
					chErr <- err
					return
				}
				g.Tick()
				bg.Fill(bg.Bounds(), white, screen.Src)
				data, err := g.Render()
				if err != nil {
					chErr <- err
					return
				}
				app.drawGrid(data, bg)
				w.Copy(pos0, bg, bg.Bounds(), screen.Src, nil)
				bg.Release()
				w.Send(paint.Event{})

				// Smooth Framerate
				deltaT := (time.Now().UTC().UnixNano() - tStart) / 10e6
				sleep := app.tick - int(deltaT)
				if sleep > 0 {
					log.Printf("Sleeping %d millis", sleep)
					time.Sleep(time.Duration(sleep) * time.Millisecond)
				} else {
					log.Printf("Overdue %d millis", -sleep)					
				}
			}
		}()

		// Loop Screen Thread until interrupted
		for !interrupt {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					interrupt = true
					chInterrupt <- true
				}
			case paint.Event:
				w.Publish()
			}
		}
	})

	// Wait for interruption or error signal
	for {
		select {
			case err := <- chErr:
				return err
			case <- chInterrupt:
				return nil
		}
	}
}

func (app *screenApp) getPixelRect(x, y int) image.Rectangle {
	return image.Rectangle{
		image.Point{x*app.scale, y*app.scale}, 
		image.Point{(x*app.scale)+app.scale, (y*app.scale)+app.scale }}
}

func (app *screenApp) drawGrid(grid [][]bool, output screen.Texture) {
	for iRow, row := range grid {
		for iCell, cell := range row {
			if cell { 
				output.Fill(app.getPixelRect(iCell, iRow), black, screen.Src) }
		}
	}
}