// Usage:
//  go run chapter3/category_accuracy.go chapter3/labeled.csv
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

	// Calculate the mean absolute error and mean squared error.
	if observed, err := table.UintColumn(table.Columns[0], 0); err != nil {
		log.Println(err)
		return -1
	} else if predicted, err := table.UintColumn(table.Columns[1], 0); err != nil {
		log.Println(err)
		return -1
	} else if len(observed) != len(predicted) {
		log.Println("Observed and predicted samples mismatch")
		return -1
	} else if classes, err := table.UintValues(table.Columns[0]); err != nil {
		log.Println(err)
		return -1
	} else {
		for _, class := range classes {
			var true_positive, false_positive uint
			var true_negative, false_negative uint
			for i := range observed {
				if observed[i] == class {
					if predicted[i] == observed[i] {
						true_positive++
					} else {
						false_negative++
					}
				} else {
					if predicted[i] == observed[i] {
						false_positive++
					} else {
						true_negative++
					}
				}
			}

			precision := float64(true_positive) / float64(true_positive+false_positive)
			recall := float64(true_positive) / float64(true_positive+false_negative)

			fmt.Println("Class", class, "=>")
			fmt.Printf("  precision=%.2f\n", precision)
			fmt.Printf("  recall=%.2f\n", recall)
			fmt.Println("")
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
