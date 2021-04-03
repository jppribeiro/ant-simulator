package simulator

import (
	"fmt"
	"math"
	"os"

	"github.com/jppribeiro/ant-simulator/internal/ant"
	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/veandco/go-sdl2/sdl"
)

func Run(wTitle string, wWidth int32, wHeight int32) int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	ants := make([]*ant.Ant, 50)

	foragingMarkers := []*ant.Marker{}
	retrievingMarkers := []*ant.Marker{}

	for i, _ := range ants {
		ants[i] = ant.NewAnt(sdl.Point{200, 200})
	}

	window, err := sdl.CreateWindow(wTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		wWidth, wHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	running := true

	counter := 0

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.SetDrawColor(80, 255, 60, 255)

		drawCircle(sdl.Point{200, 200}, 15, renderer)

		renderer.SetDrawColor(255, 255, 255, 255)

		renderer.SetDrawColor(255, 255, 255, 255)

		for _, a := range ants {
			if counter%3 == 0 {
				switch a.CurrentState {
				case ant.Foraging:
					foragingMarkers = append(foragingMarkers, a.PlaceMarker())
				case ant.Retrieving:
					retrievingMarkers = append(retrievingMarkers, a.PlaceMarker())
				}
			}

			a.Move()
			renderer.DrawPoint(int32(a.Pos.X)-1, int32(a.Pos.Y))
			renderer.DrawPoint(int32(a.Pos.X)+1, int32(a.Pos.Y))
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y)-1)
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y)+1)
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y))
		}

		delForaging := []int{}

		for i, f := range foragingMarkers {
			a := 1 - (config.ANT_CONFIG.FORAGING_TTL-f.TTL)/config.ANT_CONFIG.FORAGING_TTL

			renderer.SetDrawColor(uint8(50*a), uint8(80*a), uint8(255*a), 255)
			f.TTL--
			renderer.DrawPoint(int32(f.Pos.X), int32(f.Pos.Y))
			if f.TTL <= 0 {
				delForaging = append(delForaging, i)
			}
		}

		for _, ref := range delForaging {
			foragingMarkers[len(foragingMarkers)-1], foragingMarkers[ref] = foragingMarkers[ref], foragingMarkers[len(foragingMarkers)-1]
			foragingMarkers = foragingMarkers[:len(foragingMarkers)-1]
		}

		renderer.Present()
		sdl.Delay(33)
		counter++
	}

	return 0
}

func drawCircle(pos sdl.Point, radius int32, w *sdl.Renderer) {
	for i := -radius; i < radius; i++ {
		for j := -radius; j < radius; j++ {
			if math.Pow(float64(j), 2)+math.Pow(float64(i), 2) <= float64(radius*radius) {
				w.DrawPoint(pos.X-j, pos.Y-i)
			}
		}
	}
}
