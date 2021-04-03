package vector

import "math"

type Vector struct {
	X float32
	Y float32
}

func New(x float32, y float32) Vector {
	v := Vector{x, y}

	return v
}

func (v *Vector) Scale(m float32) {
	v.X *= m
	v.Y *= m
}

func (v *Vector) Unit() {
	m := Magnitude(*v)

	v.X /= m
	v.Y /= m
}

func (v *Vector) MirrorX() {
	v.X = -v.X
}

func (v *Vector) MirrorY() {
	v.Y = -v.Y
}

func (v *Vector) Mirror() {
	v.MirrorX()
	v.MirrorY()
}

func Magnitude(v Vector) float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)))
}
