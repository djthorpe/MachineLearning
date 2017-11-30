// Usage:
//  go get -u gonum.org/v1/gonum/stat
//  go run chisquare_01.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	// Define observed and expected values. Most
	// of the time these will come from your
	// data (website visits, etc.).
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Println("Observed", observed)
	fmt.Println("Expected", expected)
	fmt.Println("chiSquare", chiSquare)

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
