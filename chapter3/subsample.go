// Usage:
//  go run chapter3/subsample.go chapter3/time_series.csv
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	// Utilities for reading data
	"github.com/djthorpe/MachineLearning/util"
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

	// One in every four is the testing set
	training_rows := make([]int, 0)
	testing_rows := make([]int, 0)
	for row := 0; row < len(table.Rows); row++ {
		if row%4 == 0 {
			testing_rows = append(testing_rows, row)
		} else {
			training_rows = append(training_rows, row)
		}
	}
	if training_set, err := table.Subsample(training_rows); err != nil {
		log.Println("Unable to subsample training set:", err)
		return -1
	} else if testing_set, err := table.Subsample(testing_rows); err != nil {
		log.Println("Unable to subsample testing set:", err)
		return -1
	} else {
		fmt.Println("Sample size =", len(table.Rows))
		fmt.Println("Training set size =", len(training_set.Rows))
		fmt.Println("Testing set size =", len(testing_set.Rows))
	}

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
