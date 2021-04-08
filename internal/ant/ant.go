package ant

import (
	"math"
	"math/rand"
	"time"

	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/jppribeiro/ant-simulator/internal/marker"
	"github.com/jppribeiro/ant-simulator/internal/vector"
	"github.com/veandco/go-sdl2/sdl"
)

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
	SourceDecay  float64
}

func NewAnt(o sdl.Point) *Ant {
	pos := vector.Vector{
		X: float64(o.X),
		Y: float64(o.Y),
	}

	rand.Seed(time.Now().UnixNano())

	v := rand.Int31n(config.ANT_CONFIG.MAX_V-config.ANT_CONFIG.MIN_V) + config.ANT_CONFIG.MIN_V

	rand.Seed(time.Now().UnixNano())
	dR := rand.Float64() * 2 * math.Pi

	d := vector.New(math.Sin(dR), math.Cos(dR))

	c := rand.Int31n(6) + 4

	ant := Ant{
		Pos:          &pos,
		Vel:          &v,
		Dir:          &d,
		Clock:        c,
		CurrentState: Foraging,
		SourceDecay:  config.WORLD_CONFIG.ANT_SOURCE_DECAY,
	}

	return &ant
}

func (a *Ant) wobble() {
	rand.Seed(time.Now().UnixNano())
	beta := (rand.Float64() - 0.5) / 1.5

	/*
		beta := (math.Sin(20*float64(counter)/50) +
			math.Sin(2*float64(counter)/80) +
			math.Sin(2*float64(counter)/4)) / 48
	*/
	rMatrix := vector.NewRotationMatrix2D(beta)

	a.Dir.Rotate(rMatrix)
}

func (a *Ant) placeMarker() *marker.Marker {
	pos := *a.Pos

	return marker.NewMarker(pos, a.SourceDecay)
}

func (a *Ant) adjustDirection(counter int, markers *[]*marker.Marker) {
	//target := vector.Scale(*a.Dir, config.ANT_CONFIG.VIEW_D)
	//var c float32 = 0

	intensity := 0.0
	target := *a.Dir

	for _, m := range *markers {
		mVector := vector.Subtract(*m.Pos, *a.Pos)

		if vector.Magnitude(mVector) > config.ANT_CONFIG.VIEW_D {
			continue
		}

		if vector.UnitAngle(vector.Unit(mVector), *a.Dir) >= config.ANT_CONFIG.VIEW_ANGLE {
			continue
		}

		intensity += m.Intensity
		target.Add(vector.Scale(mVector, m.Intensity))
	}

	target.Unit()

	*a.Dir = target
	a.wobble()
}

/***********************
* 	PUBLIC FUNCTIONS   *
************************/

func (a *Ant) SetState(s state) {
	a.Dir.Mirror()
	a.CurrentState = s
}

func (a *Ant) RefreshMarkerSource() {
	a.SourceDecay = config.WORLD_CONFIG.ANT_SOURCE_DECAY
}

func (a *Ant) ResolveMarker(counter int, foragingMarkers *[]*marker.Marker, retrievingMarkers *[]*marker.Marker) {
	if counter%int(a.Clock) != 0 || a.SourceDecay <= 0 {
		return
	}

	switch a.CurrentState {
	case Foraging:
		*foragingMarkers = append(*foragingMarkers, a.placeMarker())
	case Retrieving:
		*retrievingMarkers = append(*retrievingMarkers, a.placeMarker())
	}
}

func (a *Ant) Move(counter int, foragingMarkers *[]*marker.Marker, retrievingMarkers *[]*marker.Marker) {
	switch a.CurrentState {
	case Foraging:
		a.adjustDirection(counter, retrievingMarkers)
	case Retrieving:
		a.adjustDirection(counter, foragingMarkers)
	}

	if a.Pos.X+a.Dir.X*float64(*a.Vel) >= float64(config.WINDOW.X) || a.Pos.X+a.Dir.X*float64(*a.Vel) <= 0 {
		a.Dir.MirrorX()
	}

	if a.Pos.Y+a.Dir.Y*float64(*a.Vel) >= float64(config.WINDOW.Y) || a.Pos.Y+a.Dir.Y*float64(*a.Vel) <= 0 {
		a.Dir.MirrorY()
	}

	a.Pos.X += a.Dir.X * float64(*a.Vel)
	a.Pos.Y += a.Dir.Y * float64(*a.Vel)

	a.SourceDecay -= 3
}
