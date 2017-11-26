// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_01.go iris.csv
//  open iris.csv_hist.png
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

///////////////////////////////////////////////////////////////////////////////

func ParseFloat(string_value string) (float64, error) {
	// Check for empty value
	if len(strings.TrimSpace(string_value)) == 0 {
		return 0.0, ErrEmpty
	}
	// Convert to integer
	if float_value, err := strconv.ParseFloat(string_value, 64); err != nil {
		return 0.0, err
	} else {
		return float64(float_value), nil
	}
}

func ReadPetalLength(rows [][]string) ([]float64, error) {
	petal_length := make([]float64, 0)
	for line_number, row := range rows {
		if len(row) != 5 {
			return nil, fmt.Errorf("Line %v: Expected 5 values", line_number+1)
		}
		// Skip header
		if *flagSkipHeader && line_number == 0 {
			continue
		}
		// Retrieve petal_length
		if petal_length_value, err := ParseFloat(row[2]); err != nil {
			return nil, err
		} else {
			petal_length = append(petal_length, petal_length_value)
		}
	}

	return petal_length, nil
}

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}
	if f, err := os.Open(flag.Arg(0)); err != nil {
		log.Println(err)
		return -1
	} else {
		defer f.Close()
		r := csv.NewReader(f)
		if rows, err := r.ReadAll(); err != nil {
			log.Println(err)
			return -1
		} else if petal_length, err := ReadPetalLength(rows); err != nil {
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
			if err := p.Save(4*vg.Inch, 4*vg.Inch, flag.Arg(0)+"_hist.png"); err != nil {
				log.Println(err)
				return -1
			}
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
