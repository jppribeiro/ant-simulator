package vector

import "math"

type Matrix2D struct {
	i float64
	j float64
	k float64
	l float64
}

func NewMatrix2D(i float64, j float64, k float64, l float64) Matrix2D {
	return Matrix2D{i, j, k, l}
}

func NewRotationMatrix2D(beta float64) Matrix2D {
	return Matrix2D{
		i: math.Cos(beta),
		j: -math.Sin(beta),
		k: math.Sin(beta),
		l: math.Cos(beta),
	}
}
