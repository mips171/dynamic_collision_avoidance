package main

import "math"

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v Vector) Subtract(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

func (v Vector) Multiply(factor float64) Vector {
	return Vector{v.X * factor, v.Y * factor}
}

func (v Vector) Divide(factor float64) Vector {
	return Vector{v.X / factor, v.Y / factor}
}

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector{v.X / magnitude, v.Y / magnitude}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Distance(other Vector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}
