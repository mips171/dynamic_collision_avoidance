package dubins

import (
	"math"
	"testing"
)

func TestComputeLSL(t *testing.T) {
	tests := []struct {
		name           string
		start          State
		end            State
		turningRadius  float64
		expectedPath   Path
		expectedError  error
	}{
		{
			name: "Simple Left Turn",
			start: State{X: 0, Y: 0, Theta: 0},
			end: State{X: 1, Y: 1, Theta: math.Pi / 2},
			turningRadius: 1,
			expectedPath: Path{
				Start: State{X: 0, Y: 0, Theta: 0},
				End: State{X: 1, Y: 1, Theta: math.Pi / 2},
				Segments: [3]SegmentType{Left, Straight, Left},
				Lengths: [3]float64{2.356194490192345, 1.1102230246251565e-16, 0.7853981633974483},
				TurningRadius: 1,
			},
			expectedError: nil,
		},
		{
			name: "Simple Right Turn",
			start: State{X: 0, Y: 0, Theta: 0},
			end: State{X: 1, Y: -1, Theta: -math.Pi / 2},
			turningRadius: 1,
			expectedPath: Path{
				Start: State{X: 0, Y: 0, Theta: 0},
				End: State{X: 1, Y: -1, Theta: -math.Pi / 2},
				Segments: [3]SegmentType{Left, Straight, Left},
				Lengths: [3]float64{0.7853981633974483, 1.1102230246251565e-16, 2.356194490192345},
				TurningRadius: 1,
			},
			expectedError: nil,
		},

		{
			name: "Large Turning Radius",
			start: State{X: 0, Y: 0, Theta: 0},
			end: State{X: 10, Y: 10, Theta: math.Pi},
			turningRadius: 5,
			expectedPath: Path{
				Start: State{X: 0, Y: 0, Theta: 0},
				End: State{X: 10, Y: 10, Theta: math.Pi},
				Segments: [3]SegmentType{Left, Straight, Left},
				Lengths: [3]float64{11.780972450961723, 14.142135623730955, 27.48893571891069},
				TurningRadius: 5,
			},
			expectedError: nil,
		},

		{
			name: "Tight U-Turn",
			start: State{X: 0, Y: 0, Theta: 0},
			end: State{X: 0, Y: 0, Theta: math.Pi},
			turningRadius: 1,
			expectedPath: Path{
				Start: State{X: 0, Y: 0, Theta: 0},
				End: State{X: 0, Y: 0, Theta: math.Pi},
				Segments: [3]SegmentType{Left, Straight, Left},
				Lengths: [3]float64{1.5707963267948966, 0, 4.71238898038469},
				TurningRadius: 1,
			},
			expectedError: nil,
		},

		{
			name: "Straight Line Path",
			start: State{X: 0, Y: 0, Theta: 0},
			end: State{X: 5, Y: 0, Theta: 0},
			turningRadius: 1,
			expectedPath: Path{
				Start: State{X: 0, Y: 0, Theta: 0},
				End: State{X: 5, Y: 0, Theta: 0},
				Segments: [3]SegmentType{Left, Straight, Left},
				Lengths: [3]float64{1.5707963267948966, 3, 1.5707963267948966},
				TurningRadius: 1,
			},
			expectedError: nil,
		},

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path, err := ComputeLSL(test.start, test.end, test.turningRadius)

			if err != test.expectedError {
				t.Errorf("Expected error: %v, but got: %v", test.expectedError, err)
			}

			if path != test.expectedPath {
				t.Errorf("Expected path: %v, but got: %v", test.expectedPath, path)
			}
		})
	}
}