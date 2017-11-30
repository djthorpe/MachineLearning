// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_01.go iris.csv
//  open iris.csv_hist.png
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/djthorpe/MachineLearning/util"
)

///////////////////////////////////////////////////////////////////////////////

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	table, _ := util.NewTable()
	filename := flag.Arg(0)
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}
	if petal_length, err := table.FloatColumn(table.Columns[2], math.NaN()); err != nil {
		log.Println(err)
		return -1
	} else if p, err := plot.New(); err != nil {
		log.Println(err)
		return -1
	} else {
		p.Title.Text = fmt.Sprintf("Histogram of petal_length")

		// Create a histogram of our values drawn from the standard normal
		h, err := plotter.NewHist(plotter.Values(petal_length), 16)
		if err != nil {
			log.Println(err)
			return -1
		}

		// Normalize the histogram.
		h.Normalize(1)

		// Add the histogram to the plot.
		p.Add(h)

		// Save the plot to a PNG file.

		if err := p.Save(4*vg.Inch, 4*vg.Inch, path.Base(filename)+"_hist.png"); err != nil {
			log.Println(err)
			return -1
		}
	}

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
