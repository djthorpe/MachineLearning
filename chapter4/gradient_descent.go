// Usage:
//  go run chapter4/gradient_descent.go chapter4/advertising.csv
package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"path"

	// Frameworks
	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	LEARNING_RATE = 0.001
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() < 1 {
		log.Println("Expected file argument", flag.NArg())
		return -1
	}

	table, _ := util.NewTable()
	filename := flag.Arg(0)
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}

	x_column := "Sales"
	y_column := "TV"
	if flag.NArg() > 1 {
		y_column = flag.Arg(1)
	}

	if x_data, err := table.FloatColumn(x_column, 0); err != nil {
		log.Println("Unable to create X samples:", err)
		return -1
	} else if y_data, err := table.FloatColumn(y_column, 0); err != nil {
		log.Println("Unable to create Y samples:", err)
		return -1
	} else if plot, err := plot.New(); err != nil {
		log.Println("Unable to create plot:", err)
		return -1
	} else {
		plot.X.Label.Text = x_column
		plot.Y.Label.Text = y_column

		if scatter, err := plotter.NewScatter(plot_points(x_data, y_data)); err != nil {
			log.Println("Unable to create plot:", err)
			return -1
		} else if line, err := plotter.NewLine(line_points(x_data, y_data)); err != nil {
			log.Println("Unable to create plot:", err)
			return -1
		} else {
			line.Color = color.RGBA{B: 255, A: 255}
			plot.Add(scatter, line)
			if err := plot.Save(4*vg.Inch, 4*vg.Inch, path.Base(filename)+"_"+x_column+"_"+y_column+".png"); err != nil {
				log.Println("Unable to create plot:", err)
				return -1
			}
		}
	}
	//

	return 0
}

// Return plot points
func plot_points(x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

// Return line points
func line_points(x, y []float64) plotter.XYs {
	b, m := gradient_descent(x, y, LEARNING_RATE, 1000)
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = m*x[i] + b
	}
	return pts
}

// Calculate squared error for samples and estimated b & m
func calculate_error(x, y []float64, b, m float64) float64 {
	var total_error float64
	for i := range x {
		total_error += math.Pow(y[i]-(m*x[i]+b), 2)
	}
	return total_error / float64(len(x))
}

// Step function for computing new values of b and m
func step(x, y []float64, b, m float64, rate float64) (float64, float64) {
	b_gradient, m_gradient := float64(0), float64(0)
	N := float64(len(x))
	for i := range x {
		b_gradient += -(2 / N) * (y[i] - (m*x[i] + b))
		m_gradient += -(2 / N) * x[i] * (y[i] - (m*x[i] + b))
	}
	return b - (rate * b_gradient), m - (rate * m_gradient)
}

// Return values of b and m
func gradient_descent(x, y []float64, rate float64, epochs uint) (float64, float64) {
	var b, m float64
	for i := uint(0); i < epochs; i++ {
		b, m = step(x, y, b, m, rate)
		fmt.Println("epoch=", i, "b=", b, "m=", m, "err=", calculate_error(x, y, b, m))

	}
	return b, m
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
