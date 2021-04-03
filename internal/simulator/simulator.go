package simulator

import (
	"fmt"
	"os"

	"github.com/jppribeiro/ant-simulator/internal/ant"
	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/veandco/go-sdl2/sdl"
)

var home *sdl.Rect
var food *sdl.Rect
var ants = make([]*ant.Ant, config.WORLD_CONFIG.ANT_NUM)
var foragingMarkers = []*ant.Marker{}
var retrievingMarkers = []*ant.Marker{}

type rgb struct {
	r int32
	g int32
	b int32
	a int32
}

func Run(wTitle string, wWidth int32, wHeight int32) int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	prepareScene()

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
		renderer.DrawRect(home)

		renderer.SetDrawColor(230, 255, 60, 255)
		renderer.DrawRect(food)

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
			evalState(a)

			renderer.DrawPoint(int32(a.Pos.X)-1, int32(a.Pos.Y))
			renderer.DrawPoint(int32(a.Pos.X)+1, int32(a.Pos.Y))
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y)-1)
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y)+1)
			renderer.DrawPoint(int32(a.Pos.X), int32(a.Pos.Y))
		}

		resolveMarkers(&foragingMarkers, renderer, rgb{50, 80, 255, 255})
		resolveMarkers(&retrievingMarkers, renderer, rgb{230, 255, 80, 255})

		renderer.Present()
		sdl.Delay(33)
		counter++
	}

	return 0
}

func resolveMarkers(markers *[]*ant.Marker, renderer *sdl.Renderer, color rgb) {
	i := 0

	for _, f := range *markers {
		a := 1 - (config.ANT_CONFIG.MARKER_TTL-f.TTL)/config.ANT_CONFIG.MARKER_TTL
		renderer.SetDrawColor(uint8(float32(color.r)*a), uint8(float32(color.g)*a), uint8(float32(color.b)*a), uint8(color.a))
		f.TTL--
		renderer.DrawPoint(int32(f.Pos.X), int32(f.Pos.Y))
		if f.TTL > 0 {
			(*markers)[i] = f
			i++
		}
	}

	for j := i; j < len(*markers); j++ {
		(*markers)[j] = nil
	}

	(*markers) = (*markers)[:i]
}

func prepareScene() {
	for i := range ants {
		ants[i] = ant.NewAnt(sdl.Point{
			X: int32(config.WORLD_CONFIG.HOME_POS_X + config.WORLD_CONFIG.HOME_SIZE/2),
			Y: int32(config.WORLD_CONFIG.HOME_POS_Y + config.WORLD_CONFIG.HOME_SIZE/2),
		})
	}

	home = &sdl.Rect{
		X: int32(config.WORLD_CONFIG.HOME_POS_X),
		Y: int32(config.WORLD_CONFIG.HOME_POS_X),
		W: int32(config.WORLD_CONFIG.HOME_SIZE),
		H: int32(config.WORLD_CONFIG.HOME_SIZE),
	}

	food = &sdl.Rect{
		X: int32(config.WORLD_CONFIG.FOOD_POS_X),
		Y: int32(config.WORLD_CONFIG.FOOD_POS_Y),
		W: int32(config.WORLD_CONFIG.FOOD_SIZE),
		H: int32(config.WORLD_CONFIG.FOOD_SIZE),
	}
}

func evalState(a *ant.Ant) {
	switch a.CurrentState {
	case ant.Foraging:
		if foundFood(a) {
			a.SetState(ant.Retrieving)
		}
	}
}

func foundFood(a *ant.Ant) bool {
	if a.Pos.X > config.WORLD_CONFIG.FOOD_POS_X &&
		a.Pos.X < config.WORLD_CONFIG.FOOD_POS_X+config.WORLD_CONFIG.FOOD_SIZE &&
		a.Pos.Y > config.WORLD_CONFIG.FOOD_POS_Y &&
		a.Pos.Y < config.WORLD_CONFIG.FOOD_POS_Y+config.WORLD_CONFIG.FOOD_SIZE {
		return true
	}

	return false
}
