// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run mean_squared_error.go time_series.csv
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	// Utilities for reading data
	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

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

	// Calculate the mean absolute error and mean squared error.
	if observed, err := table.FloatColumn(table.Columns[0], math.NaN()); err != nil {
		log.Println(err)
		return -1
	} else if predicted, err := table.FloatColumn(table.Columns[1], math.NaN()); err != nil {
		log.Println(err)
		return -1
	} else if len(observed) != len(predicted) {
		log.Println("Observed and predicted samples mismatch")
		return -1
	} else {
		var mAE float64
		var mSE float64
		for i := range observed {
			mAE += math.Abs(observed[i]-predicted[i]) / float64(len(observed))
			mSE += math.Pow(observed[i]-predicted[i], 2) / float64(len(observed))
		}
		// Output the MAE and MSE value to standard out.
		fmt.Printf("MAE = %0.2f\n", mAE)
		fmt.Printf("MSE = %0.2f\n", mSE)

		// Calculate R squared
		fmt.Printf("R^2 = %0.2f\n", stat.RSquaredFrom(observed, predicted, nil))

	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
