// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_02.go iris.csv
//  open iris.csv_hist.png
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/djthorpe/MachineLearning/util"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", false, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	table, _ := util.NewTable()
	if err := table.ReadCSV(flag.Arg(0), *flagSkipHeader, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}

	fmt.Println(table)

	if column, err := table.FloatColumn(table.Columns[3], math.NaN()); err != nil {
		log.Println(err)
		return -1
	} else {
		fmt.Println(column)
	}

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
