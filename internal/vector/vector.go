package vector

import "math"

type Vector struct {
	X float64
	Y float64
}

func New(x float64, y float64) Vector {
	v := Vector{x, y}

	return v
}

func (v *Vector) Scale(m float64) {
	v.X *= m
	v.Y *= m
}

func Scale(v Vector, m float64) Vector {
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

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vector) Rotate(m Matrix2D) {
	v.X = m.i*v.X + m.j*v.Y
	v.Y = m.k*v.X + m.l*v.Y
}

func Subtract(v1 Vector, v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

func Magnitude(v Vector) float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func Dot(v1 Vector, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Unit(v Vector) Vector {
	m := Magnitude(v)

	return Vector{v.X / m, v.Y / m}
}

// UnitAngle determines the angle between two unit vectors
// It saves computation time because we avoid having to calculate vector magnitude
func UnitAngle(v1 Vector, v2 Vector) float64 {
	return math.Acos(float64(Dot(v1, v2)))
}
