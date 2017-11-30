// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_02.go iris.csv
//  open iris.csv_boxplots.png
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
	if p, err := plot.New(); err != nil {
		log.Println(err)
		return -1
	} else {
		p.Title.Text = fmt.Sprintf("Box plots")
		p.Y.Label.Text = fmt.Sprintf("Values")

		// Create the box for our data
		w := vg.Points(50)

		// Create a box plot for each of the feature columns in the dataset.
		for i, name := range table.Columns {
			// If the column is one of the feature columns, let's create
			// a histogram of the values.
			if name != "Name" {
				// Create a plotter.Values
				if v, err := table.FloatColumn(name, math.NaN()); err != nil {
					log.Println(err)
					return -1
				} else if b, err := plotter.NewBoxPlot(w, float64(i), plotter.Values(v)); err != nil {
					log.Println(err)
					return -1
				} else {
					p.Add(b)
				}
			}
		}

		// Set the X axis of the plot to nominal with
		// the given names for x=0, x=1, etc.
		p.NominalX("SepalLength", "SepalWidth", "PetalLength", "PetalWidth")

		if err := p.Save(4*vg.Inch, 4*vg.Inch, path.Base(filename)+"_boxplots.png"); err != nil {
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
