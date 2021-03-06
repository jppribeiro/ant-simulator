package simulator

import (
	"fmt"
	"os"

	"github.com/jppribeiro/ant-simulator/internal/ant"
	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/jppribeiro/ant-simulator/internal/marker"
	"github.com/veandco/go-sdl2/sdl"
)

var home *sdl.Rect
var food *sdl.Rect

var homeStore int32
var foodStore int32

var ants = make([]*ant.Ant, config.WORLD_CONFIG.ANT_NUM)
var foragingMarkers = []*marker.Marker{}
var retrievingMarkers = []*marker.Marker{}

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

		renderer.SetDrawColor(80, 155, 60, 255)
		renderer.DrawRect(home)

		renderer.SetDrawColor(255, 155, 60, 255)
		renderer.DrawRect(food)

		renderer.SetDrawColor(180, 155, 80, 255)

		for _, a := range ants {
			a.Move(counter, foragingMarkers, retrievingMarkers)

			a.ResolveMarker(counter, &foragingMarkers, &retrievingMarkers)

			evalWorldState(a)

			renderer.DrawRect(&sdl.Rect{
				X: int32(a.Pos.X),
				Y: int32(a.Pos.Y),
				W: 2,
				H: 2,
			})
		}

		resolveMarkers(&foragingMarkers, renderer, rgb{120, 255, 255, 255})
		resolveMarkers(&retrievingMarkers, renderer, rgb{255, 190, 150, 255})

		renderer.Present()
		sdl.Delay(33)
		counter++
	}

	return 0
}

func resolveMarkers(markers *[]*marker.Marker, renderer *sdl.Renderer, color rgb) {
	i := 0

	for _, m := range *markers {
		a := 1 - (config.ANT_CONFIG.MARKER_TTL*config.WORLD_CONFIG.ANT_SOURCE_DECAY-m.TTL*m.DecayConst)/(config.ANT_CONFIG.MARKER_TTL*config.WORLD_CONFIG.ANT_SOURCE_DECAY)

		renderer.SetDrawColor(uint8(float64(color.r)*a), uint8(float64(color.g)*a), uint8(float64(color.b)*a), uint8(color.a))

		m.Decay()

		renderer.DrawPoint(int32(m.Pos.X), int32(m.Pos.Y))

		if m.TTL > 0 {
			(*markers)[i] = m
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

	homeStore = 0

	food = &sdl.Rect{
		X: int32(config.WORLD_CONFIG.FOOD_POS_X),
		Y: int32(config.WORLD_CONFIG.FOOD_POS_Y),
		W: int32(config.WORLD_CONFIG.FOOD_SIZE),
		H: int32(config.WORLD_CONFIG.FOOD_SIZE),
	}

	foodStore = 100
}

func evalWorldState(a *ant.Ant) {
	if foundHome(a) && a.CurrentState == ant.Retrieving {
		homeStore++
		a.SetState(ant.Foraging)
	}

	if foundFood(a) && a.CurrentState == ant.Foraging {
		foodStore--
		a.SetState(ant.Retrieving)
	}
}

func foundFood(a *ant.Ant) bool {
	if a.Pos.X > config.WORLD_CONFIG.FOOD_POS_X &&
		a.Pos.X < config.WORLD_CONFIG.FOOD_POS_X+config.WORLD_CONFIG.FOOD_SIZE &&
		a.Pos.Y > config.WORLD_CONFIG.FOOD_POS_Y &&
		a.Pos.Y < config.WORLD_CONFIG.FOOD_POS_Y+config.WORLD_CONFIG.FOOD_SIZE &&
		foodStore > 0 {

		a.RefreshMarkerSource()

		return true
	}

	return false
}

func foundHome(a *ant.Ant) bool {
	if a.Pos.X > config.WORLD_CONFIG.HOME_POS_X &&
		a.Pos.X < config.WORLD_CONFIG.HOME_POS_X+config.WORLD_CONFIG.HOME_SIZE &&
		a.Pos.Y > config.WORLD_CONFIG.HOME_POS_Y &&
		a.Pos.Y < config.WORLD_CONFIG.HOME_POS_Y+config.WORLD_CONFIG.HOME_SIZE {

		a.RefreshMarkerSource()

		return true
	}

	return false
}
