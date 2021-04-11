package vector

import (
	"math"
	"reflect"
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
		want float64
	}{
		{
			name: "Expect 90º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{0, 1}},
			want: math.Pi / 2,
		},
		{
			name: "Expect 45º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{math.Sqrt(2) / 2, math.Sqrt(2)}},
			want: math.Pi / 4,
		},
		{
			name: "Expect 135º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-math.Sqrt(2) / 2, math.Sqrt(2)}},
			want: 3 * math.Pi / 4,
		},
		{
			name: "Expect 180º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-1, 0}},
			want: math.Pi,
		},
		{
			name: "Expect 135º Angle",
			args: args{v1: Vector{1, 0}, v2: Vector{-math.Sqrt(2) / 2, -math.Sqrt(2)}},
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
		want float64
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
		X float64
		Y float64
	}

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Vector [1, 0]",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: 1,
		},
		{
			name: "Vector [0, 1]",
			fields: fields{
				X: 0,
				Y: 1,
			},
			want: 1,
		},
		{
			name: "Vector [0, 0]",
			fields: fields{
				X: 0,
				Y: 0,
			},
			want: 0,
		},
		{
			name: "Vector [-1, 0]",
			fields: fields{
				X: -1,
				Y: 0,
			},
			want: 1,
		},
		{
			name: "Vector [0, -1]",
			fields: fields{
				X: 0,
				Y: -1,
			},
			want: 1,
		},
		{
			name: "Vector [3, 4]",
			fields: fields{
				X: 3,
				Y: 4,
			},
			want: 5,
		},
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
		want float64
	}{
		{
			name: "Vector [1, 0]",
			args: args{
				v: Vector{
					X: 1,
					Y: 0,
				},
			},
			want: 1,
		},
		{
			name: "Vector [0, 1]",
			args: args{
				v: Vector{
					X: 0,
					Y: 1,
				},
			},
			want: 1,
		},
		{
			name: "Vector [0, 0]",
			args: args{
				v: Vector{
					X: 0,
					Y: 0,
				},
			},
			want: 0,
		},
		{
			name: "Vector [-1, 0]",
			args: args{
				v: Vector{
					X: -1,
					Y: 0,
				},
			},
			want: 1,
		},
		{
			name: "Vector [0, -1]",
			args: args{
				v: Vector{
					X: 0,
					Y: -1,
				},
			},
			want: 1,
		},
		{
			name: "Vector [3, 4]",
			args: args{
				v: Vector{
					X: 3,
					Y: 4,
				},
			},
			want: 5,
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

func TestNew(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "Expect valid vector",
			args: args{2, 2},
			want: Vector{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Scale(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		m float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector
	}{
		{
			name:   "Expect self to be scaled",
			fields: fields{1, 1},
			args:   args{2},
			want:   Vector{2, 2},
		},
		{
			name:   "Expect self to be scaled to zero",
			fields: fields{1, 1},
			args:   args{0},
			want:   Vector{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			v.Scale(tt.args.m)

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("New() = %v, want %v", v, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		v Vector
		m float64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "Expect self to be scaled",
			args: args{
				v: Vector{1, 1},
				m: 2,
			},
			want: Vector{2, 2},
		},
		{
			name: "Expect self to be scaled to zero",
			args: args{
				v: Vector{1, 1},
				m: 0,
			},
			want: Vector{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scale(tt.args.v, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Unit(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector
	}{
		{
			name:   "Expect [2, 0] -> [1, 0]",
			fields: fields{2, 0},
			want:   Vector{1, 0},
		},
		{
			name:   "Expect [2, 2] -> [sqr 2 / 2, sqr 2 / 2]",
			fields: fields{2, 2},
			want: Vector{
				X: math.Floor((math.Sqrt(2)/2)*10000) / 10000,
				Y: math.Floor((math.Sqrt(2)/2)*10000) / 10000,
			},
		},
		{
			name:   "Expect [-2, 2] -> [-sqr 2 / 2, sqr 2 / 2]",
			fields: fields{-2, 2},
			want: Vector{
				X: math.Floor((-math.Sqrt(2)/2)*10000) / 10000,
				Y: math.Floor((math.Sqrt(2)/2)*10000) / 10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			v.Unit()

			vFloor := Vector{
				X: math.Floor(v.X*10000) / 10000,
				Y: math.Floor(v.Y*10000) / 10000,
			}

			if !reflect.DeepEqual(vFloor, tt.want) {
				t.Errorf("Unit() = %v, want %v", vFloor, tt.want)
			}
		})
	}
}

func TestVector_MirrorX(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector
	}{
		{
			name:   "Expect [1, 0] -> [-1, 0]",
			fields: fields{1, 0},
			want:   Vector{-1, 0},
		},
		{
			name:   "Expect [0, 1] -> [0, 1]",
			fields: fields{0, 1},
			want:   Vector{0, 1},
		},
		{
			name:   "Expect [1, 1] -> [-1, 1]",
			fields: fields{1, 1},
			want:   Vector{-1, 1},
		},
		{
			name:   "Expect [-1, 1] -> [1, 1]",
			fields: fields{-1, 1},
			want:   Vector{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			v.MirrorX()

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("MirrorX() = %v, want %v", *v, tt.want)
			}
		})
	}
}
