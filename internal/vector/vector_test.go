package vector

import (
	"math"
	"testing"
)

func TestUnitAngle(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Expect 90º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{0, 1}},
			want: math.Pi / 2,
		},
		{
			name: "Expect 45º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{float32(math.Sqrt(2)) / 2, float32(math.Sqrt(2))}},
			want: math.Pi / 4,
		},
		{
			name: "Expect 135º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-float32(math.Sqrt(2)) / 2, float32(math.Sqrt(2))}},
			want: 3 * math.Pi / 4,
		},
		{
			name: "Expect 180º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-1, 0}},
			want: math.Pi,
		},
		{
			name: "Expect 135º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-float32(math.Sqrt(2)) / 2, -float32(math.Sqrt(2))}},
			want: 3 * math.Pi / 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitAngle(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("UnitAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Expect 0",
			args: args{v1: Vector{1, 0}, v2: Vector{0, 1}},
			want: 0,
		},
		{
			name: "Expect 24",
			args: args{v1: Vector{3, 4}, v2: Vector{4, 3}},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Magnitude(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Magnitude(); got != tt.want {
				t.Errorf("Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	type args struct {
		v Vector
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "??",
			args: args{Vector{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Magnitude(tt.args.v); got != tt.want {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
