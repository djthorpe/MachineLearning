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

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	//filename := flag.Arg(0)

	table, _ := util.NewTable()

	fmt.Println(table)
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
