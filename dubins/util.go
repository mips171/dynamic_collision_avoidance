package dubins

import "math"

// Utility function to normalize an angle to the range [0, 2*Pi)
func normalizeAngle(angle float64) float64 {
	for angle < 0 {
		angle += 2 * math.Pi
	}
	for angle >= 2*math.Pi {
		angle -= 2 * math.Pi
	}
	return angle
}
