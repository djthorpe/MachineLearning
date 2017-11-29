// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_02.go iris.csv
//  open iris.csv_hist.png
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/djthorpe/MachineLearning/util"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	filename := flag.Arg(0)
	table, _ := util.NewTable()
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println(err)
		return -1
	}

	fmt.Println(table)

	if column, err := table.StringColumn(table.Columns[2], "<nil>"); err != nil {
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
