package dubins

import (
	"math"
)

type State struct {
	X, Y  float64 // Position coordinates
	Theta float64 // Orientation angle in radians
}

type SegmentType int

const (
	Left SegmentType = iota
	Straight
	Right
)

type Path struct {
	Start, End    State
	Segments      [3]SegmentType
	Lengths       [3]float64
	TurningRadius float64
}

func NewState(x, y, theta float64) State {
	return State{X: x, Y: y, Theta: theta}
}

// Utility functions like normalization of angles, distance calculation etc. go here
// ComputeLSL computes the Left-Straight-Left Dubins path
func ComputeLSL(start, end State, turningRadius float64) (Path, error) {

	dx := end.X - start.X
	dy := end.Y - start.Y
	D := math.Sqrt(dx*dx+dy*dy) / turningRadius

	theta := normalizeAngle(math.Atan2(dy, dx))
	alpha := normalizeAngle(start.Theta - theta)
	beta := normalizeAngle(end.Theta - theta)

	segment1 := normalizeAngle((math.Pi / 2) - alpha)
	segment3 := normalizeAngle((math.Pi / 2) - beta)
	segment2 := D - math.Sin(segment1) - math.Sin(segment3)

	return Path{
		Start:         start,
		End:           end,
		Segments:      [3]SegmentType{Left, Straight, Left},
		Lengths:       [3]float64{segment1 * turningRadius, segment2 * turningRadius, segment3 * turningRadius},
		TurningRadius: turningRadius,
	}, nil
}

func ShortestPath(start, end State, turningRadius float64) (Path, error) {
	// Call each ComputeXYZ function
	// Compare the paths
	// Return the shortest one
	return Path{}, nil
}

func (p *Path) Sample(t float64) State {
	// Implement the logic to sample a point at distance 't' along the path
	// This will involve determining which segment the point falls on and then computing the exact position and orientation
	return State{}
}
