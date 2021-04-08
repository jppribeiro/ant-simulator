package marker

import (
	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/jppribeiro/ant-simulator/internal/vector"
)

type Marker struct {
	Pos *vector.Vector

	Intensity  float64
	TTL        float64
	DecayConst float64
}

func NewMarker(pos vector.Vector, decay float64) *Marker {
	m := Marker{
		Pos: &pos,

		Intensity:  config.WORLD_CONFIG.MARKER_TTL * decay,
		TTL:        config.WORLD_CONFIG.MARKER_TTL,
		DecayConst: decay,
	}

	return &m
}

func (m *Marker) Decay() {
	m.TTL = m.TTL - config.WORLD_CONFIG.MARKER_TTL/m.DecayConst
	m.Intensity = m.TTL * m.DecayConst
}
