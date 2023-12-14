package main

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		y        float64
		expected Vector
	}{
		{
			name:     "Positive values",
			x:        2,
			y:        3,
			expected: Vector{X: 2, Y: 3},
		},
		{
			name:     "Negative values",
			x:        -2,
			y:        -3,
			expected: Vector{X: -2, Y: -3},
		},
		{
			name:     "Zero values",
			x:        0,
			y:        0,
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NewVector(test.x, test.y)
			if result != test.expected {
				t.Errorf("NewVector incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector
		v2       Vector
		expected Vector
	}{
		{
			name:     "Positive values",
			v1:       Vector{X: 1, Y: 2},
			v2:       Vector{X: 3, Y: 4},
			expected: Vector{X: 4, Y: 6},
		},
		{
			name:     "Negative values",
			v1:       Vector{X: -1, Y: -2},
			v2:       Vector{X: -3, Y: -4},
			expected: Vector{X: -4, Y: -6},
		},
		{
			name:     "Zero values",
			v1:       Vector{X: 0, Y: 0},
			v2:       Vector{X: 0, Y: 0},
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v1.Add(test.v2)
			if result != test.expected {
				t.Errorf("Addition incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector
		v2       Vector
		expected Vector
	}{
		{
			name:     "Positive values",
			v1:       Vector{X: 5, Y: 6},
			v2:       Vector{X: 3, Y: 4},
			expected: Vector{X: 2, Y: 2},
		},
		{
			name:     "Negative values",
			v1:       Vector{X: -5, Y: -6},
			v2:       Vector{X: -3, Y: -4},
			expected: Vector{X: -2, Y: -2},
		},
		{
			name:     "Zero values",
			v1:       Vector{X: 0, Y: 0},
			v2:       Vector{X: 0, Y: 0},
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v1.Subtract(test.v2)
			if result != test.expected {
				t.Errorf("Subtraction incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		v        Vector
		factor   float64
		expected Vector
	}{
		{
			name:     "Positive values",
			v:        Vector{X: 2, Y: 3},
			factor:   2.5,
			expected: Vector{X: 5, Y: 7.5},
		},
		{
			name:     "Negative values",
			v:        Vector{X: -2, Y: -3},
			factor:   -2.5,
			expected: Vector{X: 5, Y: 7.5},
		},
		{
			name:     "Zero values",
			v:        Vector{X: 0, Y: 0},
			factor:   2.5,
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v.Multiply(test.factor)
			if result != test.expected {
				t.Errorf("Multiplication incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector
		v2       Vector
		expected float64
	}{
		{
			name:     "Positive values",
			v1:       Vector{X: 0, Y: 0},
			v2:       Vector{X: 3, Y: 4},
			expected: 5.0,
		},
		{
			name:     "Negative values",
			v1:       Vector{X: 0, Y: 0},
			v2:       Vector{X: -3, Y: -4},
			expected: 5.0,
		},
		{
			name:     "Zero values",
			v1:       Vector{X: 0, Y: 0},
			v2:       Vector{X: 0, Y: 0},
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v1.Distance(test.v2)
			if math.Abs(result-test.expected) > 1e-6 {
				t.Errorf("Distance calculation incorrect. Expected %f, got %f", test.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		v        Vector
		factor   float64
		expected Vector
	}{
		{
			name:     "Positive values",
			v:        Vector{X: 6, Y: 8},
			factor:   2,
			expected: Vector{X: 3, Y: 4},
		},
		{
			name:     "Negative values",
			v:        Vector{X: -6, Y: -8},
			factor:   -2,
			expected: Vector{X: 3, Y: 4},
		},
		{
			name:     "Zero values",
			v:        Vector{X: 0, Y: 0},
			factor:   2,
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v.Divide(test.factor)
			if result != test.expected {
				t.Errorf("Division incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		v        Vector
		expected Vector
	}{
		{
			name:     "Positive values",
			v:        Vector{X: 3, Y: 4},
			expected: Vector{X: 0.6, Y: 0.8},
		},
		{
			name:     "Negative values",
			v:        Vector{X: -3, Y: -4},
			expected: Vector{X: -0.6, Y: -0.8},
		},
		{
			name:     "Zero values",
			v:        Vector{X: 0, Y: 0},
			expected: Vector{X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v.Normalize()
			if math.Abs(result.X-test.expected.X) > 1e-6 || math.Abs(result.Y-test.expected.Y) > 1e-6 {
				t.Errorf("Normalization incorrect. Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		v        Vector
		expected float64
	}{
		{
			name:     "Positive values",
			v:        Vector{X: 3, Y: 4},
			expected: 5.0,
		},
		{
			name:     "Negative values",
			v:        Vector{X: -3, Y: -4},
			expected: 5.0,
		},
		{
			name:     "Zero values",
			v:        Vector{X: 0, Y: 0},
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.v.Magnitude()
			if math.Abs(result-test.expected) > 1e-6 {
				t.Errorf("Magnitude calculation incorrect. Expected %f, got %f", test.expected, result)
			}
		})
	}
}
