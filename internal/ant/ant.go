package ant

import (
	"math"
	"math/rand"
	"time"

	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/jppribeiro/ant-simulator/internal/vector"
	"github.com/veandco/go-sdl2/sdl"
)

func NewAnt(o sdl.Point) *Ant {
	pos := vector.Vector{
		X: float32(o.X),
		Y: float32(o.Y),
	}

	rand.Seed(time.Now().UnixNano())

	v := rand.Int31n(config.ANT_CONFIG.MAX_V-config.ANT_CONFIG.MIN_V) + config.ANT_CONFIG.MIN_V

	rand.Seed(time.Now().UnixNano())
	dR := rand.Float64() * 2 * math.Pi

	d := vector.New(float32(math.Sin(dR)), float32(math.Cos(dR)))

	c := rand.Int31n(10) + 7

	ant := Ant{
		Pos:          &pos,
		Vel:          &v,
		Dir:          &d,
		Clock:        c,
		CurrentState: Foraging,
	}

	return &ant
}

type state string

const (
	Foraging   state = "foraging"
	Retrieving state = "retrieving"
)

type Ant struct {
	Pos          *vector.Vector
	Vel          *int32         // Px per frame
	Dir          *vector.Vector // Random unit vector
	Clock        int32
	CurrentState state
}

type Marker struct {
	Pos *vector.Vector
	Dir *vector.Vector
	TTL float32
}

func (a *Ant) PlaceMarker() *Marker {
	pos := *a.Pos
	dir := *a.Dir
	dir.Mirror()

	m := Marker{
		Pos: &pos,
		Dir: &dir,
		TTL: config.ANT_CONFIG.MARKER_TTL,
	}

	return &m
}

func (a *Ant) Move(foragingMarkers *[]*Marker, retrievingMarkers *[]*Marker) {
	switch a.CurrentState {
	case Foraging:
		a.adjustDirection(retrievingMarkers)
	case Retrieving:
		a.adjustDirection(foragingMarkers)
	}

	if a.Pos.X+a.Dir.X*float32(*a.Vel) >= float32(config.WINDOW.X) || a.Pos.X+a.Dir.X*float32(*a.Vel) <= 0 {
		a.Dir.MirrorX()
	}

	if a.Pos.Y+a.Dir.Y*float32(*a.Vel) >= float32(config.WINDOW.Y) || a.Pos.Y+a.Dir.Y*float32(*a.Vel) <= 0 {
		a.Dir.MirrorY()
	}

	a.Pos.X += a.Dir.X * float32(*a.Vel)
	a.Pos.Y += a.Dir.Y * float32(*a.Vel)
}

func (a *Ant) ResolveMarker(counter int, foragingMarkers *[]*Marker, retrievingMarkers *[]*Marker) {
	if counter%int(a.Clock) != 0 {
		return
	}

	switch a.CurrentState {
	case Foraging:
		*foragingMarkers = append(*foragingMarkers, a.PlaceMarker())
	case Retrieving:
		*retrievingMarkers = append(*retrievingMarkers, a.PlaceMarker())
	}

}

func (a *Ant) SetState(s state) {
	a.CurrentState = s
}

func (a *Ant) adjustDirection(markers *[]*Marker) {
	target := vector.Scale(*a.Dir, config.ANT_CONFIG.VIEW_D)
	var c float32 = 1000

	for _, m := range *markers {
		mVector := vector.Subtract(*m.Pos, *a.Pos)

		if vector.Magnitude(mVector) > float32(config.ANT_CONFIG.VIEW_D) {
			continue
		}

		if vector.UnitAngle(vector.Unit(mVector), *a.Dir) >= float32(config.ANT_CONFIG.VIEW_ANGLE) {
			continue
		}

		if m.TTL < c {
			target = mVector
		}
	}

	a.Dir.Add(target)
	a.Dir.Unit()
}
