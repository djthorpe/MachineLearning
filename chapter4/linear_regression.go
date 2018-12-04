// Usage:
//  go run chapter4/linear_regression.go chapter4/data.csv
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

	if description, err := table.Describe(); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	} else {
		fmt.Println(description)
	}

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
