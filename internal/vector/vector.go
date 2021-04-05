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

func Scale(v Vector, m float32) Vector {
	return Vector{v.X * m, v.Y * m}
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

func (v *Vector) Add(v1 Vector) {
	v.X += v1.X
	v.Y += v1.Y
}

func (v *Vector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func Subtract(v1 Vector, v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

func Magnitude(v Vector) float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func Dot(v1 Vector, v2 Vector) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Unit(v Vector) Vector {
	m := Magnitude(v)

	return Vector{v.X / m, v.Y / m}
}

// UnitAngle determines the angle between two unit vectors
// It saves computation time because we avoid having to calculate vector magnitude
func UnitAngle(v1 Vector, v2 Vector) float32 {
	return float32(math.Acos(float64(Dot(v1, v2))))
}
