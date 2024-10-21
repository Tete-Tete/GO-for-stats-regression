// main.go
package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Data points from Anscombe's quartet
var datasets = []struct {
	x []float64
	y []float64
}{
	{
		x: []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
		y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	},
	{
		x: []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
		y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
	},
	{
		x: []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0},
		y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	},
	{
		x: []float64{8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 19.0, 8.0, 8.0, 8.0},
		y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	},
}

func main() {
	runLinearRegressionTests()
	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln()
}

func runLinearRegressionTests() {
	const tolerance = 1e-2

	// Perform linear regression for each dataset and measure execution time
	for i, data := range datasets {
		// Check if all x values are the same
		allXEqual := true
		for j := 1; j < len(data.x); j++ {
			if data.x[j] != data.x[0] {
				allXEqual = false
				break
			}
		}

		if allXEqual {
			fmt.Printf("Dataset %d: ERROR: Because all x values are the same, we are not able to perform linear regression\n\n", i+1)
			continue
		}

		start := time.Now()
		slope, intercept, err := calculateRegression(data.x, data.y)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf("Dataset %d: ERROR: %v\n\n", i+1, err)
			continue
		}

		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("Slope: %.6f\n", slope)
		fmt.Printf("Intercept: %.6f\n", intercept)
		fmt.Printf("Execution time: %v\n\n", duration)

		plotDataWithRegressionLine(i+1, data.x, data.y, slope, intercept)
	}
}

func calculateRegression(x, y []float64) (float64, float64, error) {
	// Convert x and y to stats.Series
	series := toSeries(x, y)

	// Calculate linear regression manually
	sum := [5]float64{}
	for i := range series {
		sum[0] += series[i].X
		sum[1] += series[i].Y
		sum[2] += series[i].X * series[i].X
		sum[3] += series[i].X * series[i].Y
		sum[4] += series[i].Y * series[i].Y
	}
	f := float64(len(series))
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)

	return gradient, intercept, nil
}

func toSeries(x, y []float64) stats.Series {
	s := make(stats.Series, len(x))
	for i := range x {
		s[i] = stats.Coordinate{X: x[i], Y: y[i]}
	}
	return s
}

func plotDataWithRegressionLine(datasetNumber int, x, y []float64, slope, intercept float64) {
	p := plot.New()

	p.Title.Text = fmt.Sprintf("Dataset %d", datasetNumber)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Create scatter plot for the data points
	pts := make(plotter.XYs, len(x))
	for i := range x {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	// Create line plot for the regression line
	line := plotter.NewFunction(func(x float64) float64 {
		return slope*x + intercept
	})
	line.Color = plotter.DefaultLineStyle.Color
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 6*vg.Inch, fmt.Sprintf("dataset_%d.png", datasetNumber)); err != nil {
		panic(err)
	}
}
