package main

import (
	"math"
	"testing"
	"time"
)

func TestLinearRegression(t *testing.T) {
	const tolerance = 1e-2
	tests := []struct {
		x         []float64
		y         []float64
		slope     float64
		intercept float64
	}{
		{
			x:         []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
			y:         []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
			slope:     0.5,
			intercept: 3.0,
		},
		{
			x:         []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
			y:         []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
			slope:     0.5,
			intercept: 3.0,
		},
		{
			x:         []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
			y:         []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
			slope:     0.5,
			intercept: 3.0,
		},
		{
			x:         []float64{8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 19.0, 8.0, 8.0, 8.0},
			y:         []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
			slope:     0.5,
			intercept: 3.0,
		},
	}

	for _, test := range tests {
		start := time.Now()
		slope, intercept, err := calculateRegression(test.x, test.y)
		duration := time.Since(start)

		if err != nil {
			t.Errorf("Error calculating regression: %v", err)
			continue
		}

		if math.IsNaN(slope) || math.IsNaN(intercept) {
			t.Errorf("Calculation resulted in NaN values for slope or intercept")
			continue
		}

		if math.Abs(slope-test.slope) > tolerance {
			t.Errorf("Expected slope %.6f, but got %.6f", test.slope, slope)
		}

		if math.Abs(intercept-test.intercept) > tolerance {
			t.Errorf("Expected intercept %.6f, but got %.6f", test.intercept, intercept)
		}

		t.Logf("Execution time: %v", duration)
	}
}
